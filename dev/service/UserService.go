package service

import (
	"log"

	"github.com/gin-gonic/gin"

	"gol/the-basics/dev/db"
	"gol/the-basics/dev/do"
	"gol/the-basics/dev/usecase"
)

type IUserService interface {
	CreateUser(ctx *gin.Context)
}

type userService struct {
	database db.Database
}

func NewUserService(database *db.Database) IUserService {
	return userService{database: *database}
}

func (userService) CreateUser(ctx *gin.Context) {
	var request do.CreateAuthUserRequest
	ctx.BindJSON(&request)

	encryptedPassword := usecase.EncryptDataOrNil([]byte(request.Password))
	log.Printf("Encrypted password value: %s", encryptedPassword)
}
