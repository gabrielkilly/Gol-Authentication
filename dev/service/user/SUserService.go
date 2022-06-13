package user

import (
	"gol/the-basics/dev/db"
	"gol/the-basics/dev/usecase"
)

type UserService struct {
	database  *db.IDatabase
	encryptor *usecase.IEncryptor
}
