package main

import (
	"gol/the-basics/dev/di"
	"gol/the-basics/dev/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	deps := di.SetupDependencies()

	router := gin.Default()

	router.Use(gin.Logger())

	router.POST("/user", usecase.ServiceContext((*deps.UserController).CreateUser))

	router.Run((*deps.GlobalConfig).Hostname)
}
