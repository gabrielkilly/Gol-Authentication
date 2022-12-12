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

func (controller *UserController) CreateUser(ctx *gin.Context) {
	var createAuthUserRequest do.CreateAuthUserRequest
	err := ctx.BindJSON(&createAuthUserRequest)
	if err != nil {
		controller.mapResponse(
			do.EmptyResponse[do.CreateAuthUserResponse](),
			exception.SHttpException{Code: http.StatusBadRequest, Message: "[UserController] " + err.Error()},
		)(ctx)
	} else {
		controller.mapResponse(
			(*controller.userService).CreateUser(createAuthUserRequest),
		)(ctx)
	}
}
