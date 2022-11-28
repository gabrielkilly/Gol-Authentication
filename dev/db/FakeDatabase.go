package db

import (
	"gol/the-basics/dev/do"
)

type fakeDatabase struct{}

func NewFakeDatabase() IDatabase {
	return fakeDatabase{}
}

func (this fakeDatabase) CreateAuthUser(username string, encryptedPassword string) (*do.CreateAuthUserResponse, error) {
	return &do.CreateAuthUserResponse{
		Id:       "tbd", //TODOs
		Username: username,
		Password: encryptedPassword,
	}, nil
}
