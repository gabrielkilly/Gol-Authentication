package tests

import (
	"gol/the-basics/main/config"
	"testing"
)

func TestInvalidFilePath(t *testing.T) {
	globalConfig, err := config.NewGlobalConfig("../resourcess/test-apper-config.yml")
	if err == nil {
		t.Errorf("Invalid file path does not cause error: %s", globalConfig)
	}
}

func TestInvalidYmlConfiguration(t *testing.T) {
	globalConfig, err := config.NewGlobalConfig("resourcess/test-app-config-broken-contract.yml")
	if err == nil {
		t.Errorf("Invalid yml configuration does not cause error: %s", globalConfig)
	}
}

func TestGlobalConfigLoad(t *testing.T) {
	globalConfig, _ := config.NewGlobalConfig("../resources/test-app-config.yml")

	if globalConfig.DatabaseUrl != "testDatabaseUrl" ||
		globalConfig.Env != "testEnv" ||
		globalConfig.Hostname != "testHostname" {
		t.Errorf("Global config variables do not match file: %s", globalConfig)
	}
}
