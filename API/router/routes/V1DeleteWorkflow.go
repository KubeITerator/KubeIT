package routes

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"kubeIT/helpers"
)

func V1DeleteWorkflow(cHandler *helpers.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {

		var workflow, project string
		workflow = c.Query("workflow")
		if workflow == "" {
			project = c.DefaultQuery("project", "kubeit")

			err := cHandler.KubeHandler.DeleteWorkflows(project)

			if err != nil {

				log.WithFields(log.Fields{
					"stage": "router",
					"topic": "delete_workflow",
					"phase": "delete_project",
					"type":  "err",
					"err":   err.Error(),
				}).Warn("Failed to delete project")
				c.AbortWithStatusJSON(400, gin.H{"error": "Failed to query project: " + project})
				return
			}

		} else {
			err := cHandler.KubeHandler.DeleteWorkflow(workflow)

			if err != nil {
				log.WithFields(log.Fields{
					"stage": "router",
					"topic": "delete_workflow",
					"phase": "delete_workflow",
					"type":  "err",
					"err":   err.Error(),
				}).Warn("Failed to delete workflow")
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
