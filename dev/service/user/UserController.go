package user

import (
	"gol/the-basics/dev/do"
	"gol/the-basics/dev/exception"
	"gol/the-basics/dev/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:generate moq -out ../../../test/mocks/IUserControllerMock.go -pkg mocks . IUserController
type IUserController interface {
	CreateUser(ctx *gin.Context)
}

type UserController struct {
	userService *IUserService
	mapResponse usecase.ResponeMapper[do.CreateAuthUserResponse]
}

func NewUserController(
	userService *IUserService,
	responseMapper usecase.ResponeMapper[do.CreateAuthUserResponse],
) IUserController {
	return &UserController{userService: userService, mapResponse: responseMapper}
}

func (this *UserController) CreateUser(ctx *gin.Context) {
	var createAuthUserRequest do.CreateAuthUserRequest
	err := ctx.BindJSON(&createAuthUserRequest)
	if err != nil {
		this.mapResponse(
			do.EmptyResponse[do.CreateAuthUserResponse](),
			exception.SHttpException{Code: http.StatusBadRequest, Message: err.Error()},
		)(ctx)
	} else {
		this.mapResponse(
			(*this.userService).CreateUser(createAuthUserRequest),
		)(ctx)
	}
	return
}
