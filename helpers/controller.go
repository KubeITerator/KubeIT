package helpers

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"kubeIT/kubectl"
	"kubeIT/s3handler"
	"strings"
)

type Controller struct {
	configName, defaultPath string
	KubeHandler             *kubectl.KubeHandler
	S3hander                *s3handler.Api
	yp                      *YamlParser
	CurrentConfig           *ConfigMapData
}

// Constructor-ish
func (ch *Controller) Init(cfname, defaultPath string, handler *kubectl.KubeHandler, s3handler *s3handler.Api) error {
	ch.S3hander = s3handler
	ch.configName = cfname
	ch.defaultPath = defaultPath
	ch.KubeHandler = handler
	ch.yp = &YamlParser{}
	err := ch.yp.Init()

	if err != nil {
		return err
	}

	_, err = ch.GetCurrentConfig()

	return err
}

func (ch *Controller) CreateNewConfig() (err error) {

	yamlcontent, err := ioutil.ReadFile(ch.defaultPath + "/default-template.yaml")

	if err != nil {
		return err
	}

	matches, err := ch.yp.ParseYaml(string(yamlcontent))

	if err != nil {
		return err
	}

	newTemplate := Template{Name: "default", Yaml: string(yamlcontent), PParams: matches}
	ch.CurrentConfig = &ConfigMapData{
		Templates: []Template{newTemplate},
	}

	err = ch.SaveConfigMap()

	if err != nil {
		return err
	}

	return nil
}

func (ch *Controller) AddAditionalTemplate(name, yaml string) (err error) {

	matches, err := ch.yp.ParseYaml(yaml)

	if err != nil {
		return err
	}

	if ch.TemplateExists(name) {
		return errors.New("template already exists")
	}

	newTemplate := Template{Name: name, Yaml: yaml, PParams: matches}
	ch.CurrentConfig.Templates = append(ch.CurrentConfig.Templates, newTemplate)

	err = ch.SaveConfigMap()

	if err != nil {
		return err
	}

	return nil
}

func (ch *Controller) TemplateExists(name string) bool {
	for _, temp := range ch.CurrentConfig.Templates {
		if temp.Name == name {
			return true
		}
	}
	return false
}

func (ch *Controller) SaveConfigMap() error {

	convToString, err := json.Marshal(ch.CurrentConfig)
	if err != nil {
		return err
	}
	mapping := map[string]string{"data": string(convToString)}
	err = ch.KubeHandler.CreateOrUpdateConfigMap(ch.configName, mapping)

	if err != nil {
		return err
	}

	return nil
}

func (ch *Controller) LoadConfigMap() error {
	cfg, err := ch.KubeHandler.GetConfigMap(ch.configName)

	if err != nil {
		return err
	}

	var mapdata ConfigMapData

	err = json.Unmarshal([]byte(cfg["data"]), &mapdata)

	if err != nil {
		fmt.Println("Error in parsing existing config")
		err = ch.CreateNewConfig()
		if err != nil {
			return err
		}
	}

	ch.CurrentConfig = &mapdata

	return nil
}

func (ch *Controller) GetCurrentConfig() (cmdata *ConfigMapData, err error) {
	if ch.CurrentConfig != nil {
		return ch.CurrentConfig, nil
	} else {
		err = ch.LoadConfigMap()
		if k8serrors.IsNotFound(err) {
			err = ch.CreateNewConfig()
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}

	}

	return ch.CurrentConfig, nil

}

func (ch *Controller) ValidateParamsAndSubmit(params map[string]string) (wfname string, missingParams []string, err error) {

	var cTemp Template

	if params["scheme"] == "" {
		return "", []string{"scheme"}, err
	} else {

		for _, tmpl := range ch.CurrentConfig.Templates {
			if tmpl.Name == params["scheme"] {
				cTemp = tmpl
				break
			}
		}

		if cTemp.Name == "" {
			fmt.Println("war error")
			return "", nil, errors.New("unknown template")
		}
	}

	// TODO: Create generic name
	var fMappings []FinalMapping

Pploop:
	for _, pparam := range cTemp.PParams {

		for k, v := range params {
			if pparam.Category+"."+pparam.Name == k {
				fMappings = append(fMappings, FinalMapping{
					ParsedParam: pparam,
					FinalValue:  v,
				})
				continue Pploop
			}
		}

		if pparam.Default != "" {
			fMappings = append(fMappings, FinalMapping{
				ParsedParam: pparam,
				FinalValue:  pparam.Default,
			})
		} else {
			missingParams = append(missingParams, pparam.Category+"."+pparam.Name)
		}
	}

	if len(missingParams) == 0 {
		yaml := ch.BuildYaml(fMappings, cTemp.Yaml)
		fmt.Println(yaml)
		wfname, err = ch.KubeHandler.StartWorkflow(yaml)
		if err != nil {
			return "", nil, err
		}
	}

	return wfname, missingParams, err

}

func (ch *Controller) BuildYaml(fMappings []FinalMapping, inputYaml string) (outputYaml string) {

	scanner := bufio.NewScanner(strings.NewReader(inputYaml))
	linenumber := 0
LineLoop:
	for scanner.Scan() {

		for _, fMapping := range fMappings {
			if linenumber == fMapping.Line {
				outputYaml += scanner.Text()[:fMapping.Loc[0]] + fMapping.FinalValue + scanner.Text()[fMapping.Loc[1]:] + "\n"
				linenumber++
				continue LineLoop
			}
		}
		outputYaml += scanner.Text() + "\n"
		linenumber++
	}

	return outputYaml
}
