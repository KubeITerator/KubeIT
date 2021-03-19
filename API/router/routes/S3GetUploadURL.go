package routes

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"kubeIT/helpers"
)

type URLResponse struct {
	URL string `json:"url"`
}

func S3GetUploadURL(cHandler *helpers.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {

		passkey := c.Query("key")
		if passkey == "" {
			log.WithFields(log.Fields{
				"stage": "router",
				"topic": "s3_get_upload_url",
				"type":  "err",
			}).Warn("No passkey specified")
			c.AbortWithStatusJSON(400, gin.H{"error": "No Passkey specified"})
			return
		}

		url, err := cHandler.S3hander.GetPresignedURL(passkey)

		if err != nil {
			log.WithFields(log.Fields{
				"stage": "router",
				"topic": "s3_get_upload_url",
				"type":  "err",
				"err":   err.Error(),
			}).Warn("Failed to get upload URL")
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to get URL: " + err.Error()})
			return
		}

		c.JSON(200, URLResponse{URL: url})
		return

	}
}
