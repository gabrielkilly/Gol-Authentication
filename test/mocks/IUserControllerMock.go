// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"github.com/gin-gonic/gin"
	"gol/authentication/main/service/user"
	"sync"
)

// Ensure, that IUserControllerMock does implement user.IUserController.
// If this is not the case, regenerate this file with moq.
var _ user.IUserController = &IUserControllerMock{}

// IUserControllerMock is a mock implementation of user.IUserController.
//
// 	func TestSomethingThatUsesIUserController(t *testing.T) {
//
// 		// make and configure a mocked user.IUserController
// 		mockedIUserController := &IUserControllerMock{
// 			CreateUserFunc: func(ctx *gin.Context)  {
// 				panic("mock out the CreateUser method")
// 			},
// 		}
//
// 		// use mockedIUserController in code that requires user.IUserController
// 		// and then make assertions.
//
// 	}
type IUserControllerMock struct {
	// CreateUserFunc mocks the CreateUser method.
	CreateUserFunc func(ctx *gin.Context)

	// calls tracks calls to the methods.
	calls struct {
		// CreateUser holds details about calls to the CreateUser method.
		CreateUser []struct {
			// Ctx is the ctx argument value.
			Ctx *gin.Context
		}
	}
	lockCreateUser sync.RWMutex
}

// CreateUser calls CreateUserFunc.
func (mock *IUserControllerMock) CreateUser(ctx *gin.Context) {
	if mock.CreateUserFunc == nil {
		panic("IUserControllerMock.CreateUserFunc: method is nil but IUserController.CreateUser was just called")
	}
	callInfo := struct {
		Ctx *gin.Context
	}{
		Ctx: ctx,
	}
	mock.lockCreateUser.Lock()
	mock.calls.CreateUser = append(mock.calls.CreateUser, callInfo)
	mock.lockCreateUser.Unlock()
	mock.CreateUserFunc(ctx)
}

// CreateUserCalls gets all the calls that were made to CreateUser.
// Check the length with:
//     len(mockedIUserController.CreateUserCalls())
func (mock *IUserControllerMock) CreateUserCalls() []struct {
	Ctx *gin.Context
} {
	var calls []struct {
		Ctx *gin.Context
	}
	mock.lockCreateUser.RLock()
	calls = mock.calls.CreateUser
	mock.lockCreateUser.RUnlock()
	return calls
}
