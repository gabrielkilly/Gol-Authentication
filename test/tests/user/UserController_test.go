package user

import (
	"gol/the-basics/main/do"
	"gol/the-basics/main/exception"
	"gol/the-basics/main/service/user"
	"gol/the-basics/test/mocks"
	"gol/the-basics/test/util"
	"testing"
)

type controllerState int

const (
	controllerSuccess controllerState = iota
	controllerFailure
)

func TestCreateUserSuccess(t *testing.T) {

	userController, expectedResponse := getUserController(controllerSuccess)

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
	userController, _ := getUserController(controllerSuccess)

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
	userController, _ := getUserController(controllerSuccess)

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
	userController, _ := getUserController(controllerFailure)

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

func getUserController(state controllerState) (user.IUserController, do.CreateAuthUserResponse) {
	expectedResponse := do.CreateAuthUserResponse{
		Username: "testname",
		Password: "testpassword",
		Id:       "testid",
	}
	var userServiceMock user.IUserService = &mocks.IUserServiceMock{
		CreateUserFunc: func(request do.CreateAuthUserRequest) (*do.CreateAuthUserResponse, exception.IHttpException) {
			if state == controllerFailure {
				return nil, exception.SHttpException{Code: 500, Message: "Error"}
			} else {
				return &expectedResponse, nil
			}
		},
	}
	controller := user.NewUserController(&userServiceMock)
	return controller, expectedResponse
}
