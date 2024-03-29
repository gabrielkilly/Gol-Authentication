// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"gol/authentication/main/db"
	"gol/authentication/main/do"
	"sync"
)

// Ensure, that IDatabaseMock does implement db.IDatabase.
// If this is not the case, regenerate this file with moq.
var _ db.IDatabase = &IDatabaseMock{}

// IDatabaseMock is a mock implementation of db.IDatabase.
//
//	func TestSomethingThatUsesIDatabase(t *testing.T) {
//
//		// make and configure a mocked db.IDatabase
//		mockedIDatabase := &IDatabaseMock{
//			CreateAuthUserFunc: func(username string, encryptedPassword string) (*do.CreateAuthUserResponse, error) {
//				panic("mock out the CreateAuthUser method")
//			},
//		}
//
//		// use mockedIDatabase in code that requires db.IDatabase
//		// and then make assertions.
//
//	}
type IDatabaseMock struct {
	// CreateAuthUserFunc mocks the CreateAuthUser method.
	CreateAuthUserFunc func(username string, encryptedPassword string) (*do.CreateAuthUserResponse, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateAuthUser holds details about calls to the CreateAuthUser method.
		CreateAuthUser []struct {
			// Username is the username argument value.
			Username string
			// EncryptedPassword is the encryptedPassword argument value.
			EncryptedPassword string
		}
	}
	lockCreateAuthUser sync.RWMutex
}

// CreateAuthUser calls CreateAuthUserFunc.
func (mock *IDatabaseMock) CreateAuthUser(username string, encryptedPassword string) (*do.CreateAuthUserResponse, error) {
	if mock.CreateAuthUserFunc == nil {
		panic("IDatabaseMock.CreateAuthUserFunc: method is nil but IDatabase.CreateAuthUser was just called")
	}
	callInfo := struct {
		Username          string
		EncryptedPassword string
	}{
		Username:          username,
		EncryptedPassword: encryptedPassword,
	}
	mock.lockCreateAuthUser.Lock()
	mock.calls.CreateAuthUser = append(mock.calls.CreateAuthUser, callInfo)
	mock.lockCreateAuthUser.Unlock()
	return mock.CreateAuthUserFunc(username, encryptedPassword)
}

// CreateAuthUserCalls gets all the calls that were made to CreateAuthUser.
// Check the length with:
//
//	len(mockedIDatabase.CreateAuthUserCalls())
func (mock *IDatabaseMock) CreateAuthUserCalls() []struct {
	Username          string
	EncryptedPassword string
} {
	var calls []struct {
		Username          string
		EncryptedPassword string
	}
	mock.lockCreateAuthUser.RLock()
	calls = mock.calls.CreateAuthUser
	mock.lockCreateAuthUser.RUnlock()
	return calls
}
