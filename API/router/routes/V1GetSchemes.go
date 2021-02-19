package routes

import (
	"github.com/gin-gonic/gin"
	"kubeIT/helpers"
)

type SchemeInfo struct {
	Name       string            `json:"name"`
	Yaml       string            `json:"yaml"`
	Parameters map[string]string `json:"parameters"`
}

func V1GetSchemes(cHandler *helpers.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {

		template := c.Query("name")
		if template == "" {
			var returnNames []string

			for _, item := range cHandler.CurrentConfig.Templates {
				returnNames = append(returnNames, item.Name)
			}
			c.JSON(200, returnNames)
			return

		} else {
			for _, item := range cHandler.CurrentConfig.Templates {
				if item.Name == template {
					sInfo := SchemeInfo{
						Name:       template,
						Yaml:       item.Yaml,
						Parameters: make(map[string]string),
					}

					for _, pparam := range item.PParams {
						sInfo.Parameters[pparam.Category+"."+pparam.Name] = pparam.Default
					}
					c.JSON(200, sInfo)
					return
				}
			}
			c.AbortWithStatusJSON(400, gin.H{"error": "Template not found: " + template})
			return
		}
	}
}
