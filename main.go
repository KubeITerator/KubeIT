package main

import (
	"fmt"
	"io/ioutil"
	"kubeIT/kubectl"
	"log"
	"time"
)

func main() {

	kH := kubectl.KubeHandler{}
	kH.StartClient("biokube")

	content, err := ioutil.ReadFile("smallex.yaml")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)

	err = kH.StartJob(text)
	if err != nil {
		fmt.Println(err.Error())
	}

	time.Sleep(10 * time.Second)

	err = kH.Delete("hello-world-test")
}
