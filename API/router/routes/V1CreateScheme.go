package routes

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"kubeIT/helpers"
)

type Template struct {
	Yaml string `json:"yaml"`
	Name string `json:"name"`
}

func V1CreateScheme(cHandler *helpers.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {

		temp := Template{}
		err := c.BindJSON(&temp)

		if err != nil {

			log.WithFields(log.Fields{
				"stage": "router",
				"topic": "create_scheme",
				"phase": "json_binding",
				"type":  "err",
				"err":   err.Error(),
			}).Warn("JSON binding failed")
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to assign template to struct: " + err.Error()})
			return
		}

		err = cHandler.AddAditionalTemplate(temp.Name, temp.Yaml)

		if err != nil {
			log.WithFields(log.Fields{
				"stage": "router",
				"topic": "apply_workflow",
				"phase": "add_scheme",
				"type":  "err",
				"err":   err.Error(),
			}).Warn("Failed to create scheme")
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to assign template to struct: " + err.Error()})
			return
		}
	}
}
