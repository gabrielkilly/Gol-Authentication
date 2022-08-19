package test

import (
	"gol/the-basics/dev/do"
	"gol/the-basics/dev/service/user"
	"gol/the-basics/test/mocks"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewUserController(t *testing.T) {
	var userServiceMock user.IUserService = &mocks.IUserServiceMock{
		CreateUserFunc: func(request do.CreateAuthUserRequest) (do.HttpResponse[do.CreateAuthUserResponse], error) {
			return do.HttpResponse[do.CreateAuthUserResponse]{}, nil
		},
	}

	userController := user.NewUserController(
		&userServiceMock, mocks.MapResponseMock[do.CreateAuthUserResponse],
	)

	if userController == nil {
		t.Error("NewUserController is returning nil")
	}
}

func TestCreateUserSuccess(t *testing.T) {
	userServiceMock := &mocks.IUserServiceMock{
		CreateUserFunc: func(request do.CreateAuthUserRequest) (do.HttpResponse[do.CreateAuthUserResponse], error) {
			return do.HttpResponse[do.CreateAuthUserResponse]{
				Code: 201,
				Data: do.CreateAuthUserResponse{
					Username: "testname",
					Password: "testpassword",
				},
			}, nil
		},
	}
	var iUserServiceMock user.IUserService = userServiceMock

	userController := user.NewUserController(
		&iUserServiceMock, mocks.MapResponseMock[do.CreateAuthUserResponse],
	)

	testRecorder := httptest.NewRecorder()
	testContext, _ := gin.CreateTestContext(testRecorder)

	userController.CreateUser(testContext)

	log.Printf("Bro hello %d", testRecorder.Result().StatusCode)
	// if mocks.ResponseMapperMock.TimesCalled != 1 {
	// 	t.Errorf("Map Response was called %d time instead of once.", mocks.ResponseMapperMock.TimesCalled)
	// }

	// if mocks.ResponseMapperMock.LastResponseCode != 201 {
	// 	t.Errorf("Http Response was %d instead of 201.", mocks.ResponseMapperMock.LastResponseCode)
	// }

}
