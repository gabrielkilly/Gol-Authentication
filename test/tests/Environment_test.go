package tests

import (
	"gol/authentication/main/config"
	"testing"
)

func TestLoadDevEnvironment(t *testing.T) {
	env, _ := config.NewEnvironment("../resources/environment/.dev-env")
	if env.GetCurrentEnv() != config.ENV_DEV {
		t.Errorf("Gettint back [%s] instead of dev", env.GetCurrentEnv())
	}
}

func TestLoadProdEnvironment(t *testing.T) {
	env, _ := config.NewEnvironment("../resources/environment/.prod-env")
	if env.GetCurrentEnv() != config.ENV_PROD {
		t.Errorf("Gettint back [%s] instead of prod", env.GetCurrentEnv())
	}
}

func TestLoadProdEnvironmentExtraSpaceAndNewLine(t *testing.T) {
	env, _ := config.NewEnvironment("../resources/environment/.prod-env-weird")
	if env.GetCurrentEnv() != config.ENV_PROD {
		t.Errorf("Gettint back [%s] instead of prod", env.GetCurrentEnv())
	}
}
