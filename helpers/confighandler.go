package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/api/errors"
	"kubeIT/kubectl"
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
