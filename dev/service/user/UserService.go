package user

import (
	"gol/the-basics/dev/db"
	"gol/the-basics/dev/do"
	"gol/the-basics/dev/usecase"
)

//go:generate moq -out ../../../test/mocks/IUserServiceMock.go -pkg mocks . IUserService
type IUserService interface {
	CreateUser(request do.CreateAuthUserRequest) (do.HttpResponse[do.CreateAuthUserResponse], error)
}

type UserService struct {
	database  *db.IDatabase
	encryptor *usecase.IEncryptor
}

func NewUserService(database *db.IDatabase, encryptor *usecase.IEncryptor) IUserService {
	return &UserService{database: database, encryptor: encryptor}
}
