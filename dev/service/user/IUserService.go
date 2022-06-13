package user

import (
	"gol/the-basics/dev/db"
	"gol/the-basics/dev/do"
	"gol/the-basics/dev/usecase"
)

type IUserService interface {
	CreateUser(request do.CreateAuthUserRequest) (do.HttpResponse[do.CreateAuthUserResponse], error)
}

func NewUserService(database *db.IDatabase, encryptor *usecase.IEncryptor) IUserService {
	return UserService{database: database, encryptor: encryptor}
}
