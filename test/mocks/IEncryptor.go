// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"gol/the-basics/dev/usecase"
	"sync"
)

// Ensure, that IEncryptorMock does implement usecase.IEncryptor.
// If this is not the case, regenerate this file with moq.
var _ usecase.IEncryptor = &IEncryptorMock{}

// IEncryptorMock is a mock implementation of usecase.IEncryptor.
//
//	func TestSomethingThatUsesIEncryptor(t *testing.T) {
//
//		// make and configure a mocked usecase.IEncryptor
//		mockedIEncryptor := &IEncryptorMock{
//			EncryptPasswordFunc: func(password string) (string, error) {
//				panic("mock out the EncryptPassword method")
//			},
//		}
//
//		// use mockedIEncryptor in code that requires usecase.IEncryptor
//		// and then make assertions.
//
//	}
type IEncryptorMock struct {
	// EncryptPasswordFunc mocks the EncryptPassword method.
	EncryptPasswordFunc func(password string) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// EncryptPassword holds details about calls to the EncryptPassword method.
		EncryptPassword []struct {
			// Password is the password argument value.
			Password string
		}
	}
	lockEncryptPassword sync.RWMutex
}

// EncryptPassword calls EncryptPasswordFunc.
func (mock *IEncryptorMock) EncryptPassword(password string) (string, error) {
	if mock.EncryptPasswordFunc == nil {
		panic("IEncryptorMock.EncryptPasswordFunc: method is nil but IEncryptor.EncryptPassword was just called")
	}
	callInfo := struct {
		Password string
	}{
		Password: password,
	}
	mock.lockEncryptPassword.Lock()
	mock.calls.EncryptPassword = append(mock.calls.EncryptPassword, callInfo)
	mock.lockEncryptPassword.Unlock()
	return mock.EncryptPasswordFunc(password)
}

// EncryptPasswordCalls gets all the calls that were made to EncryptPassword.
// Check the length with:
//
//	len(mockedIEncryptor.EncryptPasswordCalls())
func (mock *IEncryptorMock) EncryptPasswordCalls() []struct {
	Password string
} {
	var calls []struct {
		Password string
	}
	mock.lockEncryptPassword.RLock()
	calls = mock.calls.EncryptPassword
	mock.lockEncryptPassword.RUnlock()
	return calls
}