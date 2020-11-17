package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kubeIT/API/apistructs"
	"kubeIT/helpers"
	"kubeIT/kubectl"
)

func V1ApplyWorkflow(cHandler *helpers.ConfigHandler, kHandler *kubectl.KubeHandler) gin.HandlerFunc {
	return func(c *gin.Context) {

		parameters := apistructs.WorkflowParams{}
		err := c.BindJSON(&parameters)
		if err != nil {
			fmt.Println("CreateTemplate: Unknown JSON, cannot bind request to struct")
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(400, gin.H{"error": "Unknown JSON, cannot bind request to struct."})
			return
		}

		fmt.Println(parameters)
		c.JSON(200, gin.H{
			"status": "Successful",
		})
	}
}
