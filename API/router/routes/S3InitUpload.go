package routes

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"kubeIT/helpers"
)

type InitRequest struct {
	Filename string `json:"filename"`
	Multi    bool   `json:"multi"`
}
type InitResponse struct {
	Passkey string `json:"passkey"`
}

func S3InitUpload(cHandler *helpers.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {

		initRequest := InitRequest{}
		err := c.BindJSON(&initRequest)

		if err != nil {
			log.WithFields(log.Fields{
				"stage": "router",
				"topic": "s3_init_upload",
				"phase": "json_binding",
				"type":  "err",
				"err":   err.Error(),
			}).Warn("JSON binding failed")
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to assign Request to struct: " + err.Error()})
			return
		}

		passkey, err := cHandler.S3hander.InitUpload(initRequest.Filename, initRequest.Multi)

		if err != nil {
			log.WithFields(log.Fields{
				"stage": "router",
				"topic": "s3_init_upload",
				"phase": "init_upload",
				"type":  "err",
				"err":   err.Error(),
			}).Warn("S3 init failed")
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to init S3: " + err.Error()})
			return
		}

		c.JSON(200, InitResponse{Passkey: passkey})
		return

	}
}
