package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kubeIT/API/apistructs"
)

func V1ApplyWorkflow() gin.HandlerFunc {
	return func(c *gin.Context) {

		bioWorkflow := apistructs.IteratedWorkflow{}
		err := c.BindJSON(&bioWorkflow)
		fmt.Println(c.Accepted)
		if err != nil {
			fmt.Println("CreateTemplate: Unknown JSON, cannot bind request to struct")
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(400, gin.H{"error": "Unknown JSON, cannot bind request to struct."})
			return
		}
		c.JSON(200, gin.H{
			"status":    "Successful",
			"batchdata": bioWorkflow.Data.Inputs,
			"poddata":   bioWorkflow.Workflow.SplitStrategy,
		})
	}
}
