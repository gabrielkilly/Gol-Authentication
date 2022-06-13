package db

import "gol/the-basics/dev/do"

type IDatabase interface {
	CreateAuthUser(username string, encryptedPassword string) (*do.CreateAuthUserResponse, error)
}
