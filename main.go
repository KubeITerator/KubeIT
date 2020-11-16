package main

import (
	"fmt"
	"kubeIT/helpers"
	"kubeIT/kubectl"
	"os"
)

func main() {

	kH := kubectl.KubeHandler{}
	kH.StartClient("biokube")

	cH := helpers.ConfigHandler{}
	err := cH.Init("defaultconfig", "/home/beavis/go/src/kubeIT/default-settings/", &kH)

	//content, err := ioutil.ReadFile("workflowtemplate.yaml")
	//yparser := helpers.YamlParser{}
	//err = yparser.Init()
	//
	if err != nil {
		fmt.Println("Error in init")
		os.Exit(2)
	}

	ml, err := cH.GetDefaults()

	if err != nil {
		fmt.Println("Error in parsing")
		os.Exit(2)
	}

	fmt.Println(ml)

	//
	//matches, err := yparser.ParseYaml(string(content))
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//fmt.Println(matches)
	//
	//bytes, err := json.Marshal(matches)
	//
	//fmt.Println(string(bytes))

	//testsplit := strings.Split(string(content), "\n")

	//fmt.Println(len(testsplit))
	//fmt.Println(testsplit[3][16:44])

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
