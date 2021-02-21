package main

import (
	"fmt"
	network "kubeIT/API/router"
	"kubeIT/helpers"
	"kubeIT/kubectl"
	"kubeIT/s3handler"
	"os"
)

func main() {

	token := os.Getenv("TOKEN")
	s3ip := os.Getenv("S3IP")
	s3region := os.Getenv("S3REGION")
	namespace := os.Getenv("NAMESPACE")
	basebucket := os.Getenv("BASEBUCKET")

	if token == "" {
		fmt.Println("ERROR: envvar TOKEN must be specified")
		os.Exit(2)
	}

	if s3ip == "" {
		fmt.Println("ERROR: envvar S3IP must be specified")
		os.Exit(2)
	}

	if s3region == "" {
		fmt.Println("ERROR: envvar S3REGION must be specified")
		os.Exit(2)
	}

	if basebucket == "" {
		fmt.Println("ERROR: envvar BASEBUCKET must be specified")
		os.Exit(2)
	}

	if namespace == "" {
		fmt.Println("ERROR: envvar NAMESPACE must be specified")
		os.Exit(2)
	}

	kH := kubectl.KubeHandler{}
	kH.StartClient(namespace)

	s3 := s3handler.Api{}
	s3.InitS3(s3ip, s3region, basebucket)

	cH := helpers.Controller{}
	err := cH.Init("kubeit-config", "/kubeit/default-settings", &kH, &s3)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error in configHandler init")
		os.Exit(2)
	}

	router := network.Router{}
	router.Init(token)
	router.CreateRoutes(&cH)

}
