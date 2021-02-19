package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kubeIT/helpers"
)

func S3FinishUpload(cHandler *helpers.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {

		passkey := c.Query("key")
		if passkey == "" {
			fmt.Println("No Passkey specified")
			c.AbortWithStatusJSON(400, gin.H{"error": "No Passkey specified"})
			return
		}

		err := cHandler.S3hander.FinishUpload(passkey)

		if err != nil {
			fmt.Println("Failed to finish URL")
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to finish Upload: " + err.Error()})
			return
		}

		c.JSON(200, gin.H{})
		return

	}
}
