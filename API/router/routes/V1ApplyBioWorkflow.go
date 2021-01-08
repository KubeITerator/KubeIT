package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kubeIT/helpers"
)

func V1ApplyWorkflow(cHandler *helpers.ConfigHandler) gin.HandlerFunc {
	return func(c *gin.Context) {

		parameters := make(map[string]string)
		err := c.BindJSON(&parameters)
		if err != nil {
			fmt.Println("CreateTemplate: Unknown JSON, cannot bind request to struct")
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(400, gin.H{"error": "Unknown JSON, cannot bind request to struct."})
			return
		}

		wfname, missing, err := cHandler.ValidateParamsAndSubmit(parameters)

		if err != nil {
			fmt.Println("Failed Template Creation: Unknown Parameters or error")
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(400, gin.H{"error": "Unknown JSON, cannot bind request to struct."})
			return
		}

		if len(missing) > 0 {
			c.JSON(200, gin.H{
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
