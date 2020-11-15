package helpers

import (
	"encoding/json"
	"k8s.io/apimachinery/pkg/api/errors"
	"kubeIT/kubectl"
)

type ConfigHandler struct {
	handler *kubectl.KubeHandler
}

// Constructor-ish
func (ch *ConfigHandler) Init(handler *kubectl.KubeHandler) {
	ch.handler = handler
}

func (ch *ConfigHandler) GetWFConfig() {
	cfg, err := ch.handler.GetConfigMap("default-config")

	if errors.IsNotFound(err) {
		ch.SaveWFConfig()
		return
	}

}

func (ch *ConfigHandler) SaveWFConfig() {

}

func (ch *ConfigHandler) GetDefaultConfig() {

}

func (ch *ConfigHandler) SaveDefaultConfig() {

}
