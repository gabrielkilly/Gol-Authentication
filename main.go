package main

import (
	"gol/authentication/main/di"

	"github.com/gin-gonic/gin"
)

func main() {
	deps := di.SetupDependencies()

	router := gin.Default()

	router.Use(gin.Logger())

	router.POST("/user", (*deps.UserController).CreateUser)

	router.Run((*deps.GlobalConfig).Hostname)
}
