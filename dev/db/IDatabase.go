package db

import "gol/the-basics/dev/do"

//go:generate moq -out ../../test/mocks/IDatabase.go -pkg mocks . IDatabase
type IDatabase interface {
	CreateAuthUser(username string, encryptedPassword string) (*do.CreateAuthUserResponse, error)
}
