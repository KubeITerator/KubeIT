package routes

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"kubeIT/helpers"
)

func V1ApplyWorkflow(cHandler *helpers.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {

		parameters := make(map[string]string)
		err := c.BindJSON(&parameters)
		if err != nil {

			log.WithFields(log.Fields{
				"stage": "router",
				"topic": "apply_workflow",
				"phase": "json_binding",
				"type":  "err",
				"err":   err.Error(),
			}).Warn("JSON binding failed")
			c.AbortWithStatusJSON(400, gin.H{"error": "Unknown JSON, cannot bind request to struct."})
			return
		}

		wfname, missing, err := cHandler.ValidateParamsAndSubmit(parameters)

		if err != nil {
			log.WithFields(log.Fields{
				"stage": "router",
				"topic": "apply_workflow",
				"phase": "parm_validation_submit",
				"type":  "err",
				"err":   err.Error(),
			}).Warn("Failed template instantiation: Unknown parameters or error")
			c.AbortWithStatusJSON(400, gin.H{"error": "Unknown JSON, cannot bind request to struct."})
			return
		}
		if len(missing) > 0 {
			c.JSON(400, gin.H{
				"status":  "Missing parameters",
				"missing": missing,
			})
		} else {
			c.JSON(200, gin.H{
				"status": "Successful",
				"wfname": wfname,
			})
		}

	}
}
