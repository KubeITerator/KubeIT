package routes

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"kubeIT/helpers"
)

func S3GetDownloadURL(cHandler *helpers.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {

		passkey := c.Query("key")
		if passkey == "" {
			log.WithFields(log.Fields{
				"stage": "router",
				"topic": "s3_get_download_url",
				"phase": "get_key",
				"type":  "err",
			}).Warn("No passkey specified")
			c.AbortWithStatusJSON(400, gin.H{"error": "No passkey specified"})
			return
		}

		url, err := cHandler.S3hander.GetPresignedDownloadURL(passkey)

		if err != nil {

			log.WithFields(log.Fields{
				"stage": "router",
				"topic": "s3_get_download_url",
				"phase": "get_download_url",
				"type":  "err",
				"err":   err.Error(),
			}).Warn("Failed to get download URL")
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to finish upload: " + err.Error()})
			return
		}

		c.JSON(200, URLResponse{URL: url})
		return

	}
}
