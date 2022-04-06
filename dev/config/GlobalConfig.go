package config

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type GlobalConfig struct {
	Env         string `mapstructure:"env"`
	DatabaseUrl string `mapstructure:"databaseUrl"`
}

func NewGlobalConfig() GlobalConfig {
	var gc GlobalConfig

	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	loadingError := config.LoadFiles("resources/app-config.yml")
	if loadingError != nil {
		panic(loadingError)
	}

	bindingError := config.BindStruct("", &gc)
	if bindingError != nil {
		panic(bindingError)
	}

	return gc
}
