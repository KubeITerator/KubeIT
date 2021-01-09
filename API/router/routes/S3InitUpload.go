package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kubeIT/helpers"
)

type InitRequest struct {
	Key   string `json:"key"`
	Multi bool   `json:"multi"`
}
type InitResponse struct {
	Passkey string `json:"passkey"`
}

func S3InitUpload(cHandler *helpers.ConfigHandler) gin.HandlerFunc {
	return func(c *gin.Context) {

		initRequest := InitRequest{}
		err := c.BindJSON(&initRequest)

		if err != nil {
			fmt.Println("Failed assigning to json")
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to assign Request to struct: " + err.Error()})
			return
		}

		passkey, err := cHandler.S3hander.InitUpload(initRequest.Key, initRequest.Multi)

		if err != nil {
			fmt.Println("Failed to init S3")
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to init S3: " + err.Error()})
			return
		}

		c.JSON(200, InitResponse{Passkey: passkey})
		return

	}
}
