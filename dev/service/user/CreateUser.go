package user

import (
	"log"
	"net/http"

	"gol/the-basics/dev/do"
	"gol/the-basics/dev/exception"
)

func (this UserService) CreateUser(request do.CreateAuthUserRequest) (do.HttpResponse[do.CreateAuthUserResponse], error) {

	encryptedPassword, encryptionError := (*this.encryptor).EncryptPassword(request.Password)
	if encryptionError != nil {
		log.Printf("UserService.CreateUser: error encrypting password %s", encryptionError.Error())
		return do.EmptyResponse[do.CreateAuthUserResponse](), exception.SHttpException{Code: http.StatusInternalServerError, Message: encryptionError.Error()}
	}

	dbResponse, dbErr := (*this.database).CreateAuthUser(request.Username, string(encryptedPassword))
	if dbErr != nil {
		log.Printf("UserService.CreateUser: error inserting user into db %s", dbErr.Error())
		return do.EmptyResponse[do.CreateAuthUserResponse](), exception.SHttpException{Code: http.StatusInternalServerError, Message: dbErr.Error()}
	}

	return do.HttpResponse[do.CreateAuthUserResponse]{Code: http.StatusCreated, Data: *dbResponse}, nil
}
