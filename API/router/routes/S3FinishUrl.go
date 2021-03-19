package routes

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"kubeIT/helpers"
)

func S3FinishUpload(cHandler *helpers.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {

		passkey := c.Query("key")
		if passkey == "" {
			log.WithFields(log.Fields{
				"stage": "router",
				"topic": "s3_finish_upload",
			}).Warn("No passkey specified")
			c.AbortWithStatusJSON(400, gin.H{"error": "No passkey specified"})
			return
		}

		err := cHandler.S3hander.FinishUpload(passkey)

		if err != nil {

			log.WithFields(log.Fields{
				"stage": "router",
				"topic": "s3_finish_upload",
				"type":  "err",
				"err":   err.Error(),
			}).Warn("Failed to finish URL upload")
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to finish upload: " + err.Error()})
			return
		}

		c.JSON(200, gin.H{})
		return

	}
}
