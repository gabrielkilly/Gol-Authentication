package di

import (
	"gol/the-basics/dev/config"
	"gol/the-basics/dev/db"
	"gol/the-basics/dev/service"
	"gol/the-basics/dev/service/user"
	"gol/the-basics/dev/usecase"
	"log"
)

type GlobalDeps struct {
	GlobalConfig   *config.GlobalConfig
	UserController *user.IUserController
}

func SetupDependencies() GlobalDeps {
	globalConfig := killOnError(config.NewGlobalConfig("resources/app-config.yml"))
	database := db.NewDatabase(&globalConfig)
	encryptor := usecase.NewEncryptor()
	userService := service.NewUserService(&database, &encryptor)
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
