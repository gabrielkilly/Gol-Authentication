package user

import (
	"fmt"
	"gol/the-basics/dev/do"
	"gol/the-basics/dev/exception"
	"gol/the-basics/dev/usecase"

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
	bindingError := ctx.BindJSON(&createAuthUserRequest)
	if bindingError != nil {
		controller.mapResponse(
			nil,
			exception.NewInvalidParamsException(
				fmt.Sprintf("UserController.CreateUser: %s", bindingError.Error()),
			),
		)(ctx)
	} else {
		data, serviceError := (*controller.userService).CreateUser(createAuthUserRequest)
		if 
		httpResponse := do.HttpCreated(*data)
		controller.mapResponse(&httpResponse, serviceError)(ctx)
	}
}
