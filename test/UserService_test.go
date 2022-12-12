package test

import (
	"gol/the-basics/dev/do"
	"gol/the-basics/test/mocks"
	"testing"
)

type serviceState int16

const (
	dbError serviceState = iota
	encryptError
	success
)

func TestNewUserService(t *testing.T) {
	var dbMock = &mocks.IDatabaseMock{
		CreateAuthUserFunc: func(username, encryptedPassword string) (*do.CreateAuthUserResponse, error) {
			return do.EmptyResponse[do.CreateAuthUserResponse]{}, nil
		},
	}
}

func getUserService(testError )
