package main

import (
	"gol/the-basics/main/di"

	"github.com/gin-gonic/gin"
)

func main() {
	deps := di.SetupDependencies()

	router := gin.Default()

	router.Use(gin.Logger())

	router.POST("/user", (*deps.UserController).CreateUser)

	router.Run((*deps.GlobalConfig).Hostname)
}
