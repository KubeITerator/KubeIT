package routes

import (
	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"kubeIT/helpers"
)

type ArtifactResponse struct {
	Pod string `json:"pod"`
	URL string `json:"url"`
}

func V1GetResult(cHandler *helpers.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {

		workflow := c.Query("name")

		wf, err := cHandler.KubeHandler.GetWorkflow(workflow)

		if wf == nil {
			log.WithFields(log.Fields{
				"stage":    "router",
				"topic":    "get_results",
				"phase":    "find_workflow",
				"workflow": workflow,
				"type":     "missing",
			}).Warn("No workflow found")
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to query workflow: " + workflow})
			return
		}
		if err != nil {

			log.WithFields(log.Fields{
				"stage":    "router",
				"topic":    "get_results",
				"phase":    "find_workflow",
				"workflow": workflow,
				"type":     "err",
				"err":      err.Error(),
			}).Warn("Failed to query workflow")
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to query workflow: " + workflow})
			return
		}

		var status map[string]v1alpha1.NodeStatus
		status = wf.Status.Nodes

		var response []ArtifactResponse

		var afacts []v1alpha1.Artifact
		for _, stat := range status {
			if stat.Outputs != nil && stat.Outputs.Artifacts != nil {
				afacts = stat.Outputs.Artifacts
				if len(afacts) != 0 {
					for _, afec := range afacts {
						if afec.S3.Key != "" {
							url, err := cHandler.S3hander.GetPresignedDownloadInternal(afec.S3.Key)
							if err != nil {

								log.WithFields(log.Fields{
									"stage":    "router",
									"topic":    "get_results",
									"phase":    "get_artifact_url",
									"workflow": workflow,
									"type":     "err",
									"err":      err.Error(),
								}).Warn("Failed to get workflow-artifact")
								c.AbortWithStatusJSON(400, gin.H{"error": "Failed to Query S3 Key for workflow: " + workflow})
								return
							}

							response = append(response, ArtifactResponse{
								Pod: stat.DisplayName,
								URL: url,
							})
						}
					}

				}
			}
		}

		c.JSON(200, response)
		return
	}
}
