package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kubeIT/helpers"
)

func V1DeleteWorkflow(cHandler *helpers.ConfigHandler) gin.HandlerFunc {
	return func(c *gin.Context) {

		var workflow, project string
		workflow = c.Query("workflow")
		if workflow == "" {
			project = c.DefaultQuery("project", "kubeit")

			err := cHandler.KubeHandler.DeleteWorkflows(project)

			if err != nil {
				fmt.Println("Failed to query project: " + project)
				fmt.Println(err.Error())
				c.AbortWithStatusJSON(400, gin.H{"error": "Failed to query project: " + project})
				return
			}

		} else {
			err := cHandler.KubeHandler.DeleteWorkflow(workflow)

			if err != nil {
				fmt.Println("Failed to query workflow: " + workflow)
				fmt.Println(err.Error())
				c.AbortWithStatusJSON(400, gin.H{"error": "Failed to query workflow: " + workflow})
				return
			}
		}

		c.JSON(200, gin.H{
			"status": "Successful",
		})
		return

	}
}
