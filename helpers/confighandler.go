package helpers

import (
	"encoding/json"
	"io/ioutil"
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

func (ch *ConfigHandler) CreateNewConfig() error {
	//matches, err := ch.yp.ParseYaml(ch.defaultPath+"/default-template.yaml")
	//
	//if err != nil {
	//	return err
	//}

	//newConfigMapData := ConfigMapData{}
	//for _, match := range matches {
	//	combinedMapping := CombinedMappings{ParsedParam: match, }
	//}

	return nil
}

func (ch *ConfigHandler) GetConfig() ConfigMapData {
	return *ch.CurrentConfig
}

func (ch *ConfigHandler) SaveWFConfig() error {
	if ch.CurrentConfig != nil {
		convToString, err := json.Marshal(ch.CurrentConfig.mappings)
		if err != nil {
			return err
		}
		mapping := map[string]string{"yaml": ch.CurrentConfig.yaml, "mappings": string(convToString)}
		err = ch.handler.CreateOrUpdateConfigMap(ch.configName, mapping)

		if err != nil {
			return err
		}
	} else {
		ch.CreateNewConfig()
	}

	return nil
}

func (ch *ConfigHandler) GetDefaults() (ml *MappingList, err error) {
	content, err := ioutil.ReadFile(ch.defaultPath + "/mappings.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, ml)

	if err != nil {
		return nil, err
	}
	return ml, nil
}
