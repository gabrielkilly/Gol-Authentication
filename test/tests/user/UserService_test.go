package user

import (
	"errors"
	"gol/the-basics/main/db"
	"gol/the-basics/main/do"
	"gol/the-basics/main/exception"
	"gol/the-basics/main/service/user"
	"gol/the-basics/main/usecase"
	"gol/the-basics/test/mocks"
	"net/http"
	"testing"
)

type serviceState int16

const (
	serviceDbError serviceState = iota
	serviceEncryptError
	serviceSuccess
)

var defaultCreateAuthUserRequest do.CreateAuthUserRequest = do.CreateAuthUserRequest{
	Username: "testusername",
	Password: "testpassword",
}

func TestUserServiceHappyPath(t *testing.T) {
	service, expectedResponse := getUserService(serviceSuccess)

	response, _ := service.CreateUser(defaultCreateAuthUserRequest)

	if (*response).Id != (*expectedResponse).Id ||
		(*response).Username != (*expectedResponse).Username ||
		(*response).Password != (*expectedResponse).Password {
		t.Errorf("The response %s does not match expected %s", *response, *expectedResponse)
	}

}

func TestUserServiceEncryptionError(t *testing.T) {
	service, _ := getUserService(serviceEncryptError)

	_, err := service.CreateUser(defaultCreateAuthUserRequest)

	if err == nil {
		t.Error("Expected error to be thrown")
	}

	var encryptionError exception.IHttpException
	errors.As(err, &encryptionError)
	if encryptionError.StatusCode() != http.StatusInternalServerError {
		t.Errorf("Status code is %d instead of 500", encryptionError.StatusCode())
	}

}

func TestUserServiceDatabaseError(t *testing.T) {
	service, _ := getUserService(serviceDbError)

	_, err := service.CreateUser(defaultCreateAuthUserRequest)

	if err == nil {
		t.Error("Expected error to be thrown")
	}

	var encryptionError exception.IHttpException
	errors.As(err, &encryptionError)
	if encryptionError.StatusCode() != http.StatusInternalServerError {
		t.Errorf("Status code is %d instead of 500", encryptionError.StatusCode())
	}

}

func getUserService(state serviceState) (user.IUserService, *do.CreateAuthUserResponse) {
	expectedResponse := do.CreateAuthUserResponse{
		Id:       "testid",
		Username: "testusername",
		Password: "testpassword",
	}
	var mockEncryptor usecase.IEncryptor = &mocks.IEncryptorMock{
		EncryptPasswordFunc: func(password string) (string, error) {
			if state == serviceEncryptError {
				return "", errors.New("ENCRYPT ERROR!")
			} else {
				return "encryptedpassword", nil
			}
		},
	}
	var mockDb db.IDatabase = &mocks.IDatabaseMock{
		CreateAuthUserFunc: func(username, encryptedPassword string) (*do.CreateAuthUserResponse, error) {
			if state == serviceDbError {
				return nil, errors.New("DB ERROR!")
			} else {
				return &expectedResponse, nil
			}
		},
	}

	return user.NewUserService(&mockDb, &mockEncryptor), &expectedResponse
}
