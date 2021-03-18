package network

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"kubeIT/API/router/routes"
	"kubeIT/helpers"
	"math"
	"net/http"
	"os"
	"time"
)

type Router struct {
	engine *gin.Engine
	xToken string
}

func (route *Router) Init(xAuthToken string) {
	route.xToken = xAuthToken
	route.engine = gin.Default()
	route.engine.Use(Logger())
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

var timeFormat = "02/Jan/2006:15:04:05 -0700"

// Logger is the logrus logger handler
func Logger(notLogged ...string) gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknow"
	}

	var skip map[string]struct{}

	if length := len(notLogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, p := range notLogged {
			skip[p] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		if _, ok := skip[path]; ok {
			return
		}

		entry := log.WithFields(logrus.Fields{
			"statusCode": statusCode,
			"latency":    latency, // time to process
			"clientIP":   clientIP,
			"method":     c.Request.Method,
			"path":       path,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("%s - %s [%s] \"%s %s\" %d %d \"%s\" \"%s\" (%dms)", clientIP, hostname, time.Now().Format(timeFormat), c.Request.Method, path, statusCode, dataLength, referer, clientUserAgent, latency)
			if statusCode >= http.StatusInternalServerError {
				entry.Error(msg)
			} else if statusCode >= http.StatusBadRequest {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}
