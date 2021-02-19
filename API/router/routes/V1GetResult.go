package routes

import (
	"fmt"
	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/gin-gonic/gin"
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

		if wf == nil || err != nil {
			fmt.Println("Failed to query workflow: " + workflow)
			if err != nil {
				fmt.Println(err.Error())
			}
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
								fmt.Println("Failed to query workflow: " + workflow)
								fmt.Println(err.Error())
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
