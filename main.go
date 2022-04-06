package main

import (
	"gol/the-basics/dev/config"
	"gol/the-basics/dev/db"
	"gol/the-basics/dev/service"

	"github.com/gin-gonic/gin"
)

func main() {
	globalConfig := config.NewGlobalConfig()
	database := db.NewDatabase(&globalConfig)
	userService := service.NewUserService(&database)

	router := gin.Default()
	router.Use(gin.Logger())
	router.POST("/user", userService.CreateUser)
	router.Run("localhost:8080")
}
