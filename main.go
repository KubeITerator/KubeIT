package main

import (
	"fmt"
	"io/ioutil"
	"kubeIT/helpers"
	"os"
	"strings"
)

func main() {

	//kH := kubectl.KubeHandler{}
	//kH.StartClient("biokube")

	content, err := ioutil.ReadFile("workflowtemplate.yaml")
	yparser := helpers.YamlParser{}
	err = yparser.Init()

	if err != nil {
		fmt.Println("Error in parsing")
		os.Exit(2)
	}

	matches := yparser.ParseYaml(string(content))

	for _, match := range matches {

		fmt.Println(match)
	}

	testsplit := strings.Split(string(content), "\n")

	fmt.Println(len(testsplit))
	fmt.Println(testsplit[3][16:44])

	//content, err := ioutil.ReadFile("workflowtemplate.yaml")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//
	//testdata := map[string]string{"yaml":string(content)}
	//
	//err = kH.CreateOrUpdateConfigMap("testconfig", testdata)

	//cfgMap, err := kH.GetConfigMap("testconfig")
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//fmt.Println(cfgMap["yaml"])
	//text := string(content)
	//err = kH.StartJob(text)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//time.Sleep(10 * time.Second)
	//
	//err = kH.Delete("hello-world-test")
}
