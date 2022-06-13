package usecase

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type IEncryptor interface {
	EncryptPassword(password string) (string, error)
}

type Encryptor struct {
	Encrypt func(data []byte, cost int) ([]byte, error)
}

func NewEncryptor() IEncryptor {
	return Encryptor{
		Encrypt: bcrypt.GenerateFromPassword,
	}
}

func (this Encryptor) EncryptPassword(password string) (string, error) {
	hashedPassword, hashingError := this.Encrypt([]byte(password), bcrypt.MinCost)

	if hashingError != nil {
		return "", fmt.Errorf("EncryptData: Error encrypting data %w", hashingError)
	}

	return string(hashedPassword), nil
}
