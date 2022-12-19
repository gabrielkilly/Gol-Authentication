package db

import "gol/the-basics/main/do"

//go:generate moq -out ../../test/mocks/IDatabase.go -pkg mocks . IDatabase
type IDatabase interface {
	CreateAuthUser(username string, encryptedPassword string) (*do.CreateAuthUserResponse, error)
}
