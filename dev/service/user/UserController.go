package user

import (
	"gol/the-basics/dev/do"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	CreateUser(ctx *gin.Context) (do.HttpResponse[do.CreateAuthUserResponse], error)
}

type UserController struct {
	userService *IUserService
}

func NewUserController(userService *IUserService) IUserController {
	return UserController{userService: userService}
}

func (this UserController) CreateUser(ctx *gin.Context) (do.HttpResponse[do.CreateAuthUserResponse], error) {
	var createAuthUserRequest do.CreateAuthUserRequest
	ctx.BindJSON(&createAuthUserRequest)
	return (*this.userService).CreateUser(createAuthUserRequest)
}
