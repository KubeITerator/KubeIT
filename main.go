package main

import (
	"fmt"
	network "kubeIT/API/router"
	"kubeIT/helpers"
	"kubeIT/kubectl"
	"os"
)

func main() {

	kH := kubectl.KubeHandler{}
	kH.StartClient("kubeit")

	cH := helpers.ConfigHandler{}
	err := cH.Init("kubeit-defaultconfig", "/home/beavis/go/src/kubeIT/default-settings/", &kH)

	if err != nil {
		fmt.Println("Error in configHandler init")
		os.Exit(2)
	}

	router := network.Router{}
	router.Init("TEST")
	router.CreateRoutes(&cH)

}
