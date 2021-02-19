package network

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"kubeIT/API/router/routes"
	"kubeIT/helpers"
)

type Router struct {
	engine *gin.Engine
	xToken string
}

func (route *Router) Init(xAuthToken string) {
	route.xToken = xAuthToken
	route.engine = gin.Default()
	route.engine.Use(gin.Logger())
	route.engine.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Authorization", "Origin", "*"}
	route.engine.Use(cors.New(config))
}

func (route *Router) Run(address string) {
	err := route.engine.Run(address)
	if err != nil {
		panic(err)
	}
}

//noinspection ALL
func (route *Router) CreateRoutes(cHandler *helpers.Controller) {
	router := route.engine

	// Jobs / Pods Group
	v1 := router.Group("/v1")
	s3 := router.Group("/s3")

	s3.Use(route.AuthTokenMiddleware())
	{
		s3.POST("/init", routes.S3InitUpload(cHandler))
		s3.GET("/upload", routes.S3GetUploadURL(cHandler))
		s3.GET("/finish", routes.S3FinishUpload(cHandler))
		s3.GET("/download", routes.S3GetDownloadURL(cHandler))
	}

	v1.Use(route.AuthTokenMiddleware())
	{
		v1.POST("/apply", routes.V1ApplyWorkflow(cHandler))
		v1.GET("/status", routes.V1GetStatus(cHandler))
		v1.GET("/scheme", routes.V1GetSchemes(cHandler))
		v1.POST("/createscheme", routes.V1CreateScheme(cHandler))
		v1.GET("/result", routes.V1GetResult(cHandler))
		v1.GET("/delete", routes.V1DeleteWorkflow(cHandler))
	}

	router.NoRoute(route.AuthTokenMiddleware(), func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	route.Run(":9999")
}

func (route *Router) AuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		route.validateToken(c)
		c.Next()
	}
}

func (route *Router) validateToken(c *gin.Context) {
	token := c.Request.Header.Get("X-Auth-Token")

	if token == "" {
		c.AbortWithStatus(401)
	} else if token == route.xToken {
		c.Next()
	} else {
		c.AbortWithStatus(401)
	}
}
