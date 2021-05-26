package main

import (
	log "github.com/sirupsen/logrus"
	db "kubeIT/database"
	"kubeIT/server"
	"kubeIT/server/gateway"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	// LOG Formatting rules:
	// Levels: Info, Warning, Error

	// log.WithFields(log.Fields{
	//  "stage": event,
	//  "topic": topic,
	//  "key": key,
	//}).Fatal("Failed to send event")
}

func main() {

	//token := os.Getenv("TOKEN")
	//s3ip := os.Getenv("S3IP")
	//s3region := os.Getenv("S3REGION")
	//namespace := os.Getenv("NAMESPACE")
	//basebucket := os.Getenv("BASEBUCKET")
	//
	//if token == "" {
	//
	//	log.WithFields(log.Fields{
	//		"stage": "INIT",
	//		"topic": "envvars",
	//		"key":   "TOKEN",
	//	}).Fatal("Envvar TOKEN must be specified")
	//}
	//
	//if s3ip == "" {
	//	log.WithFields(log.Fields{
	//		"stage": "init",
	//		"topic": "envvars",
	//		"key":   "S3IP",
	//	}).Fatal("Envvar S3IP must be specified")
	//}
	//
	//if s3region == "" {
	//	log.WithFields(log.Fields{
	//		"stage": "init",
	//		"topic": "envvars",
	//		"key":   "S3REGION",
	//	}).Fatal("Envvar S3REGION must be specified")
	//}
	//
	//if basebucket == "" {
	//	log.WithFields(log.Fields{
	//		"stage": "init",
	//		"topic": "envvars",
	//		"key":   "BASEBUCKET",
	//	}).Fatal("Envvar BASEBUCKET must be specified")
	//}
	//
	//if namespace == "" {
	//	log.WithFields(log.Fields{
	//		"stage": "init",
	//		"topic": "envvars",
	//		"key":   "NAMESPACE",
	//	}).Fatal("Envvar NAMESPACE must be specified")
	//}
	//
	//kH := kubectl2.KubeHandler{}
	//kH.StartClient(namespace)
	//
	//s3 := s3handler2.Api{}
	//s3.InitS3(s3ip, s3region, basebucket)
	//
	//cH := helpers.Controller{}
	//err := cH.Init("kubeit-config", "/kubeit/default-settings", &kH, &s3)
	//
	//if err != nil {
	//	log.WithFields(log.Fields{
	//		"stage": "init",
	//		"topic": "confighandler",
	//		"key":   "confighandler",
	//	}).Fatal(err.Error())
	//}

	oicdClient := os.Getenv("OICDCLIENT")

	if oicdClient == "" {
		log.WithFields(log.Fields{
			"stage": "init",
			"topic": "envvars",
			"key":   "OICDCLIENT",
		}).Fatal("Envvar OICDCLIENT must be specified")
	}

	oicdSecret := os.Getenv("OICDSECRET")

	if oicdSecret == "" {
		log.WithFields(log.Fields{
			"stage": "init",
			"topic": "envvars",
			"key":   "OICDSECRET",
		}).Fatal("Envvar OICDSECRET must be specified")
	}

	database := db.Database{}

	err := database.Init()

	if err != nil {

	}

	grpc := server.Api{}
	grpc.Init(&database)

	grpcgw := gateway.Gateway{}
	grpcgw.Init(oicdClient, oicdSecret)

}
