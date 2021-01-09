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

	kH := kubectl.KubeHandler{}
	kH.StartClient("kubeit")

	s3 := s3handler.Api{}
	s3.InitS3("", "RegionOne", "kubeit")

	cH := helpers.ConfigHandler{}
	err := cH.Init("kubeit-defaultconfig", "/home/beavis/go/src/kubeIT/default-settings/", &kH, &s3)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error in configHandler init")
		os.Exit(2)
	}

	router := network.Router{}
	router.Init("TEST")
	router.CreateRoutes(&cH)

}
