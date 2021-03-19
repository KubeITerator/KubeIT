package s3handler

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type Api struct {
	S3         *s3.S3
	Session    *session.Session
	BaseBucket string
	Multiparts map[string]*MultiPart
}

type MultiPart struct {
	Parts    []*UploadedPart
	Key      string
	Multi    bool
	UploadID string
}

type UploadedPart struct {
	Etag       string
	PartNumber int64
}

func Rand16() (str string) {
	b := make([]byte, 16)
	rand.Read(b)
	str = fmt.Sprintf("%x", b)
	return
}

func (api *Api) InitS3(ip, region, basebucket string) {
	s3Config := &aws.Config{
		Endpoint: aws.String(ip),
		Region:   aws.String(region),
	}

	api.BaseBucket = basebucket

	api.Session = session.Must(session.NewSession(s3Config))

	api.Multiparts = make(map[string]*MultiPart)

	api.S3 = s3.New(api.Session)
}
func (api *Api) CheckIfExists(key string) (exists bool) {

	log.WithFields(log.Fields{
		"stage": "s3handler",
		"topic": "existence_check",
		"key":   key,
	}).Info("Checking for existence: " + key)

	input := &s3.HeadObjectInput{
		Bucket: aws.String(api.BaseBucket),
		Key:    aws.String(key),
	}

	obj, err := api.S3.HeadObject(input)
	if err != nil {
		return false
	}
	length := *obj.ContentLength
	if length == 0 {
		return false
	}
	return true
}

func (api *Api) UploadFileToS3(filename string, data []byte) (err error) {
	_, err = api.S3.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(api.BaseBucket),
		Key:                aws.String(filename),
		ACL:                aws.String("private"),
		Body:               bytes.NewReader(data),
		ContentLength:      aws.Int64(int64(len(data))),
		ContentType:        aws.String(http.DetectContentType(data)),
		ContentDisposition: aws.String("attachment"),
	})

	return err
}

func (api *Api) InitUpload(filename string, multi bool) (passkey string, err error) {

	passkey = Rand16()
	key := "/inputdata/" + passkey + "/" + filename
	if api.CheckIfExists(key) {
		return "", errors.New("key already exists")
	}

	if multi {
		multiPartOut, err := api.S3.CreateMultipartUpload(&s3.CreateMultipartUploadInput{
			Bucket: aws.String(api.BaseBucket),
			Key:    aws.String(key),
		})

		if err != nil || multiPartOut == nil {
			return "", err
		}

		api.Multiparts[passkey] = &MultiPart{
			Parts:    nil,
			Key:      key,
			Multi:    multi,
			UploadID: *multiPartOut.UploadId,
		}
	} else {
		api.Multiparts[passkey] = &MultiPart{
			Parts:    nil,
			Key:      key,
			Multi:    multi,
			UploadID: "",
		}
	}

	return passkey, nil
}

func (api *Api) GetPresignedURL(passkey string) (url string, err error) {

	part := api.Multiparts[passkey]
	if part == nil {

		return "", errors.New("unknown passkey")

	}

	if !part.Multi {
		request, _ := api.S3.PutObjectRequest(&s3.PutObjectInput{
			Bucket: aws.String(api.BaseBucket),
			Key:    aws.String(part.Key),
		})

		url, err = request.Presign(time.Minute * 180)

		return url, err

	} else {
		partNum := len(part.Parts) + 1

		req, out := api.S3.UploadPartRequest(&s3.UploadPartInput{
			Bucket:     aws.String(api.BaseBucket),
			Key:        aws.String(part.Key),
			PartNumber: aws.Int64(int64(partNum)),
			UploadId:   aws.String(part.UploadID),
		})

		err = req.Send()
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					log.WithFields(log.Fields{
						"stage": "s3handler",
						"topic": "get_presigned_url_single",
						"type":  "aerr",
						"part":  part,
						"err":   aerr.Error(),
					}).Warn("Failed to get presigned SingleURL")
					return "", aerr

				}
			} else {
				log.WithFields(log.Fields{
					"stage": "s3handler",
					"topic": "get_presigned_url_single",
					"type":  "err",
					"part":  part,
					"err":   err.Error(),
				}).Warn("Failed to get presigned SingleURL")
				return "", err
			}
		}

		if req.Error != nil {
			return "", req.Error
		}
		presignedURL, err := req.Presign(time.Minute * 180)

		part.Parts = append(part.Parts, &UploadedPart{
			Etag:       *out.ETag,
			PartNumber: int64(partNum),
		})
		api.Multiparts[passkey] = part

		return presignedURL, err

	}

}

func (api *Api) GetPresignedURLMulti(key, passkey string) (url string, e error) {

	if api.Multiparts[passkey] == nil {

		return "", errors.New("unknown passkey")

	}

	part := api.Multiparts[passkey]

	partNum := len(part.Parts) + 1

	req, out := api.S3.UploadPartRequest(&s3.UploadPartInput{
		Bucket:     aws.String(api.BaseBucket),
		Key:        aws.String(key),
		PartNumber: aws.Int64(int64(partNum)),
		UploadId:   aws.String(part.UploadID),
	})

	err := req.Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.WithFields(log.Fields{
					"stage": "s3handler",
					"topic": "get_presigned_url_multi",
					"type":  "aerr",
					"part":  part,
					"err":   aerr.Error(),
				}).Warn("Failed to get presigned MultiURL")
				return "", aerr

			}
		} else {
			log.WithFields(log.Fields{
				"stage": "s3handler",
				"topic": "get_presigned_url_multi",
				"type":  "err",
				"part":  part,
				"err":   err.Error(),
			}).Warn("Failed to get presigned MultiURL")
			return "", err
		}
	}

	if req.Error != nil {
		return "", req.Error
	}
	presignedURL, err := req.Presign(time.Minute * 180)

	part.Parts = append(part.Parts, &UploadedPart{
		Etag:       *out.ETag,
		PartNumber: int64(partNum),
	})
	api.Multiparts[passkey] = part

	return presignedURL, err
}

func (api *Api) FinishUpload(passkey string) (err error) {

	var completedPartsAWS []*s3.CompletedPart

	parts := api.Multiparts[passkey]

	if parts.Multi {
		log.WithFields(log.Fields{
			"stage": "s3handler",
			"topic": "finish_upload",
			"phase": "start",
			"part":  parts,
		}).Info("Completing multipart Upload")
		for _, part := range parts.Parts {
			awsPart := s3.CompletedPart{
				ETag:       aws.String(part.Etag),
				PartNumber: aws.Int64(part.PartNumber),
			}

			completedPartsAWS = append(completedPartsAWS, &awsPart)
			log.WithFields(log.Fields{
				"stage":  "s3handler",
				"topic":  "finish_upload_parts",
				"part":   parts,
				"etag":   *awsPart.ETag,
				"number": *awsPart.PartNumber,
			}).Debug("Part completion debug information")
		}

		_, err = api.S3.CompleteMultipartUpload(&s3.CompleteMultipartUploadInput{
			Bucket:   aws.String(api.BaseBucket),
			Key:      aws.String(parts.Key),
			UploadId: aws.String(parts.UploadID),
			MultipartUpload: &s3.CompletedMultipartUpload{
				Parts: completedPartsAWS,
			},
		})

		if err != nil {
			return err
		}

		time.Sleep(3 * time.Second)
		if api.CheckIfExists(parts.Key) {
			log.WithFields(log.Fields{
				"stage": "s3handler",
				"topic": "finish_upload",
				"phase": "end",
				"part":  parts,
			}).Info("Multipart Upload completed.")
			return nil
		} else {
			return errors.New("does not exist")
		}

	} else {

		time.Sleep(3 * time.Second)
		if api.CheckIfExists(parts.Key) {
			log.WithFields(log.Fields{
				"stage": "s3handler",
				"topic": "finish_upload",
				"phase": "end",
				"part":  parts,
			}).Info("Upload completed.")
			return nil
		} else {
			return errors.New("does not exist")
		}
	}
}

func (api *Api) DownloadFileFromS3(folder, key, path string) error {

	file, err := os.Create(path)

	defer file.Close()

	if err != nil {
		return err
	}

	downloader := s3manager.NewDownloader(api.Session)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(api.BaseBucket),
			Key:    aws.String(folder + "/" + key),
		})
	if err != nil {
		log.WithFields(log.Fields{
			"stage": "s3handler",
			"topic": "download_file",
			"type":  "err",
			"key":   key,
			"err":   err.Error(),
		}).Warn("Failed to download file from S3")
	}

	log.WithFields(log.Fields{
		"stage":    "s3handler",
		"topic":    "download_file",
		"filename": file.Name(),
		"numbytes": numBytes,
	}).Info("Download successful")

	return err
}

func (api *Api) GetPresignedDownloadURL(passkey string) (url string, err error) {

	part := api.Multiparts[passkey]
	if part == nil {

		return "", errors.New("unknown passkey")

	}

	params := &s3.GetObjectInput{
		Bucket: aws.String(api.BaseBucket),
		Key:    aws.String(part.Key),
	}

	request, _ := api.S3.GetObjectRequest(params)

	url, err = request.Presign(time.Hour * 48)

	return url, err

}

func (api *Api) GetPresignedDownloadInternal(key string) (url string, err error) {

	params := &s3.GetObjectInput{
		Bucket: aws.String(api.BaseBucket),
		Key:    aws.String(key),
	}

	request, _ := api.S3.GetObjectRequest(params)

	url, err = request.Presign(time.Hour * 48)

	return url, err

}

func (api *Api) UploadPartInternal(resp *s3.CreateMultipartUploadOutput, fileBytes []byte, partNumber, maxRetries int) (*s3.CompletedPart, error) {
	tryNum := 1
	partInput := &s3.UploadPartInput{
		Body:          bytes.NewReader(fileBytes),
		Bucket:        resp.Bucket,
		Key:           resp.Key,
		PartNumber:    aws.Int64(int64(partNumber)),
		UploadId:      resp.UploadId,
		ContentLength: aws.Int64(int64(len(fileBytes))),
	}

	for tryNum <= maxRetries {
		uploadResult, err := api.S3.UploadPart(partInput)
		if err != nil {
			if tryNum == maxRetries {
				if aerr, ok := err.(awserr.Error); ok {
					return nil, aerr
				}
				return nil, err
			}
			log.WithFields(log.Fields{
				"stage":    "s3handler",
				"topic":    "internal_part_upload",
				"filename": resp.Key,
				"partnum":  partNumber,
				"retries":  tryNum,
			}).Warn("Retrying to upload part")
			tryNum++
		} else {
			return &s3.CompletedPart{
				ETag:       uploadResult.ETag,
				PartNumber: aws.Int64(int64(partNumber)),
			}, nil
		}
	}
	return nil, nil
}

func (api *Api) CompleteMultipartUploadInternal(resp *s3.CreateMultipartUploadOutput, completedParts []*s3.CompletedPart) (*s3.CompleteMultipartUploadOutput, error) {
	completeInput := &s3.CompleteMultipartUploadInput{
		Bucket:   resp.Bucket,
		Key:      resp.Key,
		UploadId: resp.UploadId,
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: completedParts,
		},
	}
	return api.S3.CompleteMultipartUpload(completeInput)
}

func (api *Api) AbortMultipartUploadInternal(svc *s3.S3, resp *s3.CreateMultipartUploadOutput) error {
	log.WithFields(log.Fields{
		"stage":    "s3handler",
		"topic":    "abort_multipart",
		"filename": resp.Key,
		"uploadid": *resp.UploadId,
	}).Info("Aborting multipart upload")
	abortInput := &s3.AbortMultipartUploadInput{
		Bucket:   resp.Bucket,
		Key:      resp.Key,
		UploadId: resp.UploadId,
	}
	_, err := svc.AbortMultipartUpload(abortInput)
	return err
}

func (api *Api) InitMultiPartInternal(folder, key string) (resp *s3.CreateMultipartUploadOutput, err error) {
	input := &s3.CreateMultipartUploadInput{
		Bucket: aws.String(api.BaseBucket),
		Key:    aws.String(folder + "/" + key),
	}

	return api.S3.CreateMultipartUpload(input)
}
