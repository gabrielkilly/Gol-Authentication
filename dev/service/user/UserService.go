package user

import (
	"gol/the-basics/dev/db"
	"gol/the-basics/dev/do"
	"gol/the-basics/dev/exception"
	"gol/the-basics/dev/usecase"
)

//go:generate moq -out ../../../test/mocks/IUserServiceMock.go -pkg mocks . IUserService
type IUserService interface {
	CreateUser(request do.CreateAuthUserRequest) (*do.CreateAuthUserResponse, error)
}

type UserService struct {
	database  *db.IDatabase
	encryptor *usecase.IEncryptor
}

func NewUserService(database *db.IDatabase, encryptor *usecase.IEncryptor) IUserService {
	return &UserService{database: database, encryptor: encryptor}
}

func (this *UserService) CreateUser(request do.CreateAuthUserRequest) (*do.CreateAuthUserResponse, error) {

	encryptedPassword, encryptionError := (*this.encryptor).EncryptPassword(request.Password)
	if encryptionError != nil {
		return nil, exception.NewEncryptPasswordException("UserSerivce.CreateUser")
	}

	dbResponse, dbErr := (*this.database).CreateAuthUser(request.Username, string(encryptedPassword))
	if dbErr != nil {
		return nil, exception.NewDatabaseErrorException("UserService.CreateUser")
	}

	return dbResponse, nil
}
