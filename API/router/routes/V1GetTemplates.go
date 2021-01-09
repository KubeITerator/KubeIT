package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kubeIT/helpers"
)

func V1GetTemplates(cHandler *helpers.ConfigHandler) gin.HandlerFunc {
	return func(c *gin.Context) {

		template := c.Query("name")
		if template == "" {
			var returnNames []string

			err := cHandler.SaveConfigMap()
			if err != nil {
				fmt.Println("Failed to update server")
				fmt.Println(err.Error())
				c.AbortWithStatusJSON(400, gin.H{"error": "Failed to update server"})
				return
			}
			for _, item := range cHandler.CurrentConfig.Templates {
				returnNames = append(returnNames, item.Name)
			}
			c.JSON(200, returnNames)
			return

		} else {
			for _, item := range cHandler.CurrentConfig.Templates {
				if item.Name == template {
					c.JSON(200, item.Yaml)
					return
				}
			}
			c.AbortWithStatusJSON(400, gin.H{"error": "Template not found: " + template})
			return
		}
	}
}
