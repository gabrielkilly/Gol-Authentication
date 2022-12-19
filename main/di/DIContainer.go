package di

import (
	"gol/the-basics/main/config"
	"gol/the-basics/main/db"
	"gol/the-basics/main/service/user"
	"gol/the-basics/main/usecase"
	"log"
)

type GlobalDeps struct {
	GlobalConfig   *config.GlobalConfig
	UserController *user.IUserController
}

func SetupDependencies() GlobalDeps {
	globalConfig := killOnError(config.NewGlobalConfig("resources/app-config.yml"))
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
