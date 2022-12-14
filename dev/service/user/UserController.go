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
) IUserController {
	return &UserController{userService: userService}
}

func (userController *UserController) CreateUser(ctx *gin.Context) {
	var createAuthUserRequest do.CreateAuthUserRequest
	bindingError := ctx.BindJSON(&createAuthUserRequest)
	if bindingError != nil {
		usecase.MapError(
			exception.NewInvalidParamsException(
				fmt.Sprintf("UserController.CreateUser: %s", bindingError.Error()),
			),
		)(ctx)
	} else {
		data, serviceError := (*userController.userService).CreateUser(createAuthUserRequest)
		if serviceError != nil {
			usecase.MapError(serviceError)(ctx)
		} else {
			httpResponse := do.HttpCreated(*data)
			usecase.MapResponse(&httpResponse)(ctx)
		}
	}
}
