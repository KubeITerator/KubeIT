package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kubeIT/helpers"
)

type URLResponse struct {
	URL string `json:"url"`
}

func S3GetUploadURL(cHandler *helpers.ConfigHandler) gin.HandlerFunc {
	return func(c *gin.Context) {

		passkey := c.Query("key")
		if passkey == "" {
			fmt.Println("No Passkey specified")
			c.AbortWithStatusJSON(400, gin.H{"error": "No Passkey specified"})
			return
		}

		url, err := cHandler.S3hander.GetPresignedURL(passkey)

		if err != nil {
			fmt.Println("Failed to get URL")
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to get URL: " + err.Error()})
			return
		}

		c.JSON(200, URLResponse{URL: url})
		return

	}
}
