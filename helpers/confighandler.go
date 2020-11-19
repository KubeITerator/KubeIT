package helpers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/api/errors"
	"kubeIT/kubectl"
	"strings"
)

type ConfigHandler struct {
	configName, defaultPath string
	handler                 *kubectl.KubeHandler
	yp                      *YamlParser
	CurrentConfig           *ConfigMapData
}

// Constructor-ish
func (ch *ConfigHandler) Init(cfname, defaultPath string, handler *kubectl.KubeHandler) error {
	ch.configName = cfname
	ch.defaultPath = defaultPath
	ch.handler = handler
	ch.yp = &YamlParser{}
	err := ch.yp.Init()

	if err != nil {
		return err
	}

	_, err = ch.GetCurrentConfig()

	return err
}

func (ch *ConfigHandler) CreateNewConfig() (err error) {

	yamlcontent, err := ioutil.ReadFile(ch.defaultPath + "/default-template.yaml")

	if err != nil {
		return err
	}

	matches, err := ch.yp.ParseYaml(string(yamlcontent))

	if err != nil {
		return err
	}

	mappings, err := ch.GetDefaults()

	if err != nil {
		return err
	}

	var cMaps []CombinedMappings

	for _, match := range matches {
		for _, mapping := range mappings {

			if mapping.Category == match.Category {
				if mapping.Name == match.Name {
					cMaps = append(cMaps, CombinedMappings{
						ParsedParam: match,
						Defaults:    mapping.Defaults,
					})

				}

			}

		}

	}
	ch.CurrentConfig = &ConfigMapData{
		mappings: cMaps,
		yaml:     string(yamlcontent),
	}

	err = ch.SaveConfigMap()

	if err != nil {
		return err
	}

	return nil
}

func (ch *ConfigHandler) SaveConfigMap() error {

	convToString, err := json.Marshal(ch.CurrentConfig.mappings)
	if err != nil {
		return err
	}
	mapping := map[string]string{"yaml": ch.CurrentConfig.yaml, "mappings": string(convToString)}
	err = ch.handler.CreateOrUpdateConfigMap(ch.configName, mapping)

	if err != nil {
		return err
	}

	return nil
}

func (ch *ConfigHandler) LoadConfigMap() error {
	cfg, err := ch.handler.GetConfigMap(ch.configName)

	if err != nil {
		return err
	}

	var mappings []CombinedMappings

	err = json.Unmarshal([]byte(cfg["mappings"]), &mappings)

	if err != nil {
		return err
	}

	ch.CurrentConfig = &ConfigMapData{
		mappings: mappings,
		yaml:     cfg["yaml"],
	}

	return nil
}

func (ch *ConfigHandler) GetDefaults() (ml []Mapping, err error) {
	content, err := ioutil.ReadFile(ch.defaultPath + "/mappings.json")
	if err != nil {
		fmt.Println("Error in reading file")
		return nil, err
	}

	err = json.Unmarshal(content, &ml)

	if err != nil {
		fmt.Println("Error in unmarshaling")
		return nil, err
	}
	return ml, nil
}

func (ch *ConfigHandler) GetCurrentConfig() (cmdata *ConfigMapData, err error) {
	if ch.CurrentConfig != nil {
		return ch.CurrentConfig, nil
	} else {
		err = ch.LoadConfigMap()

		if errors.IsNotFound(err) {
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

func (ch *ConfigHandler) ValidateParamsAndSubmit(params map[string]interface{}) (wfname string, missingParams []string, err error) {
	ccfg, err := ch.GetCurrentConfig()

	if err != nil {
		return "", nil, err
	}

	var fMappings []FinalMapping

MapLoop:
	for _, mapping := range ccfg.mappings {
		if mapping.Required {
			for key, value := range params {
				if mapping.Category+"."+mapping.Name == key { // TODO: Reformat ugly if-conditional
					fMappings = append(fMappings, FinalMapping{
						ParsedParam: mapping.ParsedParam,
						FinalValue:  fmt.Sprintf("%v", value),
					})

					continue MapLoop

				}
			}

			missingParams = append(missingParams, mapping.Category+"."+mapping.Name)

		} else {
			fMappings = append(fMappings, FinalMapping{
				ParsedParam: mapping.ParsedParam,
				FinalValue:  mapping.Defaults.Default, // TODO: Allow for type interface values
			})
		}
	}

	if len(missingParams) == 0 {
		yaml := ch.BuildYaml(fMappings)
		fmt.Println(yaml)
		wfname, err = ch.handler.StartWorkflow(yaml)
		if err != nil {
			return "", nil, err
		}
	}

	return wfname, missingParams, err

}

func (ch *ConfigHandler) BuildYaml(fMappings []FinalMapping) (yaml string) {

	scanner := bufio.NewScanner(strings.NewReader(ch.CurrentConfig.yaml))
	linenumber := 0
LineLoop:
	for scanner.Scan() {

		for _, fMapping := range fMappings {
			if linenumber == fMapping.Line {
				yaml += scanner.Text()[:fMapping.Loc[0]] + fMapping.FinalValue + scanner.Text()[fMapping.Loc[1]:] + "\n"
				linenumber++
				continue LineLoop
			}
		}
		yaml += scanner.Text() + "\n"
		linenumber++
	}

	return yaml
}
