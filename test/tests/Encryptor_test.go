package tests

import (
	"errors"
	"gol/the-basics/main/usecase"
	"testing"
)

func TestNewEncryptor(t *testing.T) {
	encryptor := usecase.NewEncryptor()

	if encryptor == nil {
		t.Error("Problem initializing Encryptor")
	}
}

func TestPasswordEncryption(t *testing.T) {
	encryptor := usecase.NewEncryptor()

	encryptedPassword, _ := encryptor.EncryptPassword("password")

	if encryptedPassword == "password" || encryptedPassword == "" {
		t.Error("Password encryption did not take place")
	}
}

func TestPasswordEncryptionFailure(t *testing.T) {
	mockError := errors.New("Mock error")
	encryptor := usecase.Encryptor{
		Encrypt: func(data []byte, cost int) ([]byte, error) {
			return nil, mockError
		},
	}

	encryptedPassword, err := encryptor.EncryptPassword("password")

	if encryptedPassword != "" || !errors.Is(err, mockError) {
		t.Error("Password encryption failure not handled properly")
	}
}
