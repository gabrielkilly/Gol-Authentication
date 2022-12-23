package tests

import (
	golconfig "gol/authentication/main/config"
	"os"
	"testing"
)

func TestLoadDevEnvironment(t *testing.T) {
	env, _ := golconfig.NewEnvironment()
	if env.GetCurrentEnv() != golconfig.ENV_DEV {
		t.Errorf("Gettint back [%s] instead of dev", env.GetCurrentEnv())
	}
}

func TestLoadProdEnvironment(t *testing.T) {
	os.Setenv("ENV", "prod")
	env, _ := golconfig.NewEnvironment()
	if env.GetCurrentEnv() != golconfig.ENV_PROD {
		t.Errorf("Gettint back [%s] instead of prod", env.GetCurrentEnv())
	}
}
