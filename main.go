package main

import (
	"gol/the-basics/dev/config"
	"gol/the-basics/dev/service"
	"gol/the-basics/dev/util"

	"github.com/gin-gonic/gin"
)

func main() {
	util.InitializeSingletons(config.GlobalConfig)

	router := gin.Default()
	router.Use(gin.Logger())
	router.POST("/user", service.CreateUser)
	router.Run("localhost:8080")
}
