package model

import (
	"encoding/json"
	"github.com/usagiga/Nurupoga/entity"
	"io/ioutil"
)

type ConfigModelImpl struct{}

func NewConfigModel() (configModel ConfigModel) {
	return &ConfigModelImpl{}
}

func (m *ConfigModelImpl) Load(configFilePath string) (config *entity.Config, err error) {
	config = &entity.Config{}

	jsonBytes, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
