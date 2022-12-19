package config

import (
	"errors"
	"io/ioutil"
	"strings"
)

type environment struct {
	currentEnv ENV
}

type ENV string

const (
	ENV_DEV  ENV = "dev"
	ENV_PROD ENV = "prod"
)

type IEnvironment interface {
	GetCurrentEnv() ENV
	GetConfigPath() string
}

func NewEnvironment(envFilePath string) (IEnvironment, error) {
	data, readError := ioutil.ReadFile(envFilePath)

	if readError != nil {
		return nil, errors.New("Failure getting environment value")
	} else {
		parsedEnv = parseEnv(string(data))
		return environment{
			currentEnv: parseEnv(string(data)),
		}, nil
	}
}

func parseEnv(data string) (*ENV, error) {
	cleanedEnv := strings.Trim(data, " \n")
	switch cleanedEnv {
	case string(ENV_DEV):
		return &ENV_DEV, nil
	case string(ENV_PROD):
		return ENV_PROD, nil
	default: 
		return nil, nil
	}
}

func (env environment) GetCurrentEnv() ENV {
	
}
