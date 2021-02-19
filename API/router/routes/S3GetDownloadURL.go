package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kubeIT/helpers"
)

func S3GetDownloadURL(cHandler *helpers.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {

		passkey := c.Query("key")
		if passkey == "" {
			fmt.Println("No Passkey specified")
			c.AbortWithStatusJSON(400, gin.H{"error": "No Passkey specified"})
			return
		}

		url, err := cHandler.S3hander.GetPresignedDownloadURL(passkey)

		if err != nil {
			fmt.Println("Failed to finish URL")
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to finish Upload: " + err.Error()})
			return
		}

		c.JSON(200, URLResponse{URL: url})
		return

	}
}
