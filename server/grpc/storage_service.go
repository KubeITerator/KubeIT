package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	db "kubeIT/database"
	"kubeIT/pkg/grpc/common"
	"kubeIT/pkg/grpc/storage"
	"kubeIT/server/helpers"
	"kubeIT/server/s3handler"
)

type StorageService struct {
	storage.UnimplementedStorageManagementServiceServer
	database   *db.Database
	authorizer *helpers.Authorizer
	s3handler  *s3handler.Api
}

func NewStorageService(db *db.Database, authorizer *helpers.Authorizer, s3handler *s3handler.Api) *StorageService {
	return &StorageService{database: db, authorizer: authorizer, s3handler: s3handler}
}

func (sts *StorageService) InitUpload(ctx context.Context, req *storage.InitRequest) (*storage.IDmessage, error) {

	if req.StorageInfo.StorageType != storage.StorageType_S3 {
		return nil, status.Errorf(codes.Unimplemented, "InitUpload for non S3 storage types not implemented")
	}

	//stype := req.StorageInfo.TypeInfo.(*storage.StorageInfo_S3Info)

	file := storage.File{
		Id:          "",
		UserId:      req.Userid,
		Name:        req.Filename,
		StorageInfo: req.StorageInfo,
		Location:    "",
		Status:      storage.Status_INIT,
	}

	id, err := sts.database.AddStorage(&file)

	if err != nil {

	}

	return &storage.IDmessage{TransferId: id.Hex()}, nil
}
func (sts *StorageService) GetUploadUrl(context.Context, *storage.IDmessage) (*storage.UrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUploadUrl not implemented")
}
func (sts *StorageService) FinishUpload(context.Context, *storage.IDmessage) (*common.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishUpload not implemented")
}
func (sts *StorageService) DeleteUpload(context.Context, *storage.IDmessage) (*common.StatusReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUpload not implemented")
}
func (sts *StorageService) GetDownload(context.Context, *storage.IDmessage) (*storage.UrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownload not implemented")
}
