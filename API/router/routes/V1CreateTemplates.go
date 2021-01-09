package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kubeIT/helpers"
)

type Template struct {
	Yaml string `json:"yaml"`
	Name string `json:"name"`
}

func V1CreateTemplates(cHandler *helpers.ConfigHandler) gin.HandlerFunc {
	return func(c *gin.Context) {

		temp := Template{}
		err := c.BindJSON(&temp)

		if err != nil {
			fmt.Println("Failed to assign template to struct")
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to assign template to struct: " + err.Error()})
			return
		}

		err = cHandler.AddAditionalTemplate(temp.Name, temp.Yaml)

		if err != nil {
			fmt.Println("Failed to assign template to struct")
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(400, gin.H{"error": "Failed to assign template to struct: " + err.Error()})
			return
		}
	}
}
