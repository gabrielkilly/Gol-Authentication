package usecase

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncryptDataOrNil(bytes []byte) []byte {
	hashedPassword, hashingError := bcrypt.GenerateFromPassword(bytes, bcrypt.MinCost)

	if hashingError != nil {
		log.Printf("Error encrypting data: %s", hashingError.Error())
		return nil
	}

	return hashedPassword
}
