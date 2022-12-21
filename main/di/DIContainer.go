package di

import (
	"gol/authentication/main/config"
	"gol/authentication/main/db"
	"gol/authentication/main/service/user"
	"gol/authentication/main/usecase"
	"log"
)

type GlobalDeps struct {
	GlobalConfig   *config.GlobalConfig
	UserController *user.IUserController
}

func SetupDependencies() GlobalDeps {
	environment := killOnError(config.NewEnvironment(".env"))
	globalConfig := killOnError(config.NewGlobalConfig(environment.GetConfigPath()))
	database := db.NewFakeDatabase()
	encryptor := usecase.NewEncryptor()
	userService := user.NewUserService(&database, &encryptor)
	userController := user.NewUserController(&userService)

	return GlobalDeps{
		UserController: &userController,
		GlobalConfig:   &globalConfig,
	}
}

func killOnError[K any](object K, err error) K {
	if err != nil {
		log.Fatal(err.Error())
	}
	return object
}
