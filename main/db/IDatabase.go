package db

import "gol/authentication/main/do"

//go:generate moq -out ../../test/mocks/IDatabase.go -pkg mocks . IDatabase
type IDatabase interface {
	CreateAuthUser(username string, encryptedPassword string) (*do.CreateAuthUserResponse, error)
}
