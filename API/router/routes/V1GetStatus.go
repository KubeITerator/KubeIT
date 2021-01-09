package routes

import (
	"fmt"
	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/gin-gonic/gin"
	"kubeIT/helpers"
)

type WFStatus struct {
	Workflow      string `json:"workflow"`
	Status        string `json:"status"`
	Statusmessage string `json:"statusmessage"`
	Running       int    `json:"steps"`
	Finished      int    `json:"finished"`
}

func V1GetStatus(cHandler *helpers.ConfigHandler) gin.HandlerFunc {
	return func(c *gin.Context) {

		var wfStats []WFStatus
		var workflow, project string
		workflow = c.Query("workflow")
		if workflow == "" {
			project = c.DefaultQuery("project", "kubeit")

			wfs, err := cHandler.KubeHandler.GetWorkflows(project)

			if wfs == nil || err != nil {
				fmt.Println("Failed to query project: " + project)
				fmt.Println(err.Error())
				c.AbortWithStatusJSON(400, gin.H{"error": "Failed to query project: " + project})
				return
			}

			if len(wfs.Items) == 0 {
				fmt.Println("Failed to query project: " + project)
				c.AbortWithStatusJSON(400, gin.H{"error": "No workflows found for project: " + project})
				return
			}

			for _, wf := range wfs.Items {
				status := wf.Status.Nodes

				pods := 0
				finished := 0

				for _, stat := range status {
					if stat.Phase.Completed() {
						finished++
						pods++
					} else {
						pods++
					}
				}
				wfStats = append(wfStats, WFStatus{
					Workflow:      wf.Name,
					Status:        string(wf.Status.Phase),
					Statusmessage: wf.Status.Message,
					Running:       pods,
					Finished:      finished,
				})
			}

		} else {
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

			pods := 0
			finished := 0

			for _, stat := range status {
				if stat.Phase.Completed() {
					finished++
					pods++
				} else {
					pods++
				}
			}

			wfStats = append(wfStats, WFStatus{
				Workflow:      workflow,
				Status:        string(wf.Status.Phase),
				Statusmessage: wf.Status.Message,
				Running:       pods,
				Finished:      finished,
			})
		}

		c.JSON(200, wfStats)
		return

	}
}
