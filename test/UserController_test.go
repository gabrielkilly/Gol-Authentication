package test

import (
	"gol/the-basics/dev/do"
	"gol/the-basics/dev/exception"
	"gol/the-basics/dev/service/user"
	"gol/the-basics/dev/usecase"
	"gol/the-basics/test/mocks"
	"gol/the-basics/test/util"
	"testing"
)

type userControllerState int

const (
	success userControllerState = iota
	failure
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

	userController, expectedResponse := getUserController(success)

	testRecorder, testContext := util.MockGin()
	util.MockJsonPost(
		testContext,
		map[string]interface{}{"password": "testpassword", "username": "testusername"},
	)

	userController.CreateUser(testContext)

	result := testRecorder.Result()
	var bodyReponse do.CreateAuthUserResponse
	util.UnmarshalResponseBody(result.Body, &bodyReponse)

	if result.StatusCode != 201 {
		t.Errorf("Http code %d was returned instead of 201", result.StatusCode)
	}
	if bodyReponse.Password != expectedResponse.Password ||
		bodyReponse.Id != expectedResponse.Id ||
		bodyReponse.Username != expectedResponse.Username {
		t.Errorf("The response and expected response don't match [expected] %s [actual] %s", bodyReponse, expectedResponse)
	}

}

func TestCreateUserInvalidRequestNoPassword(t *testing.T) {
	userController, _ := getUserController(success)

	testRecorder, testContext := util.MockGin()

	util.MockJsonPost(
		testContext,
		map[string]interface{}{"username": "testusername"},
	)

	userController.CreateUser(testContext)

	result := testRecorder.Result()
	if result.StatusCode != 400 {
		t.Errorf("Status code is %d not 400", result.StatusCode)
	}
}

func TestCreateUserInvalidRequestNoUsername(t *testing.T) {
	userController, _ := getUserController(success)

	testRecorder, testContext := util.MockGin()

	util.MockJsonPost(
		testContext,
		map[string]interface{}{"password": "testpassword"},
	)

	userController.CreateUser(testContext)

	result := testRecorder.Result()
	if result.StatusCode != 400 {
		t.Errorf("Status code is %d not 400", result.StatusCode)
	}
}

func TestCreateUserServiceError(t *testing.T) {
	userController, _ := getUserController(failure)

	testRecorder, testContext := util.MockGin()

	util.MockJsonPost(
		testContext,
		map[string]interface{}{"password": "testpassword", "username": "testusername"},
	)

	userController.CreateUser(testContext)

	result := testRecorder.Result()
	if result.StatusCode != 500 {
		t.Errorf("Status code is %d not 500", result.StatusCode)
	}
}

func getUserController(state userControllerState) (user.IUserController, do.CreateAuthUserResponse) {
	expectedResponse := do.CreateAuthUserResponse{
		Username: "testname",
		Password: "testpassword",
		Id:       "testid",
	}
	userServiceMock := &mocks.IUserServiceMock{
		CreateUserFunc: func(request do.CreateAuthUserRequest) (do.HttpResponse[do.CreateAuthUserResponse], error) {
			if state == failure {
				return do.EmptyResponse[do.CreateAuthUserResponse](), exception.SHttpException{Code: 500, Message: "Error"}
			} else {
				return do.HttpResponse[do.CreateAuthUserResponse]{
					Code: 201, Data: expectedResponse,
				}, nil
			}
		},
	}
	var iUserServiceMock user.IUserService = userServiceMock
	controller := user.NewUserController(
		&iUserServiceMock, usecase.MapResponse[do.CreateAuthUserResponse],
	)
	return controller, expectedResponse
}
