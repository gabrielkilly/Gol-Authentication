package config

import (
	"fmt"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type GlobalConfig struct {
	Env         string `mapstructure:"env"`
	Hostname    string `mapstructure:"hostname"`
	DatabaseUrl string `mapstructure:"databaseUrl"`
}

func NewGlobalConfig(configFilePath string) (GlobalConfig, error) {
	var gc GlobalConfig

	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	loadingError := config.LoadFiles(configFilePath)
	if loadingError != nil {
		return gc, fmt.Errorf("GlobalConfig load error %s", loadingError)
	}

	bindingError := config.BindStruct("", &gc)
	if bindingError != nil {
		return gc, fmt.Errorf("GlobalConfig binding error %s", bindingError)
	}

	return gc, nil
}
