package user

import (
	"errors"
	"gol/the-basics/dev/db"
	"gol/the-basics/dev/do"
	"gol/the-basics/dev/service/user"
	"gol/the-basics/dev/usecase"
	"gol/the-basics/test/mocks"
	"testing"
)

type serviceState int16

const (
	serviceDbError serviceState = iota
	serviceEncryptError
	serviceSuccess
)

func TestUserServiceHappyPath(*testing.T) {
	// service, expectedResponse := getUserService(serviceSuccess)

	// response, _ := service.CreateUser(
	// 	do.CreateAuthUserRequest{
	// 		Username: "testusername",
	// 		Password: "testpassword",
	// 	},
	// )

	// response.

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
