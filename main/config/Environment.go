package config

import (
	"fmt"
	"os"
	"strings"
)

type environment struct {
	currentEnv ENV
}

type ENV string

const (
	ENV_DEV   ENV = "dev"
	ENV_PROD  ENV = "prod"
	env_error ENV = "error"
)

type IEnvironment interface {
	GetCurrentEnv() ENV
	GetConfigPath() string
}

func NewEnvironment() (IEnvironment, error) {
	parsedEnv, parseError := parseEnv(os.Getenv("ENV"))

	if parseError != nil {
		return nil, parseError
	}

	return environment{
		currentEnv: parsedEnv,
	}, nil
}

func parseEnv(data string) (ENV, error) {
	cleanedEnv := strings.Trim(data, " \n")
	switch cleanedEnv {
	case string(ENV_DEV):
		return ENV_DEV, nil
	case string(ENV_PROD):
		return ENV_PROD, nil
	default:
		return ENV_DEV, nil
	}
}

func (env environment) GetCurrentEnv() ENV {
	return env.currentEnv
}

func (env environment) GetConfigPath() string {
	return fmt.Sprintf("resources/vault/%s/app-config.yml", env.currentEnv)
}
