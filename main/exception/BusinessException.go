package exception

import (
	"fmt"
	"net/http"
)

type invalidParamsException struct {
	SHttpException
}

func NewInvalidParamsException(context string) invalidParamsException {
	return invalidParamsException{
		SHttpException{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("invalidParamsException fired in [%s]", context),
		},
	}
}

type encryptPasswordException struct {
	SHttpException
}

func NewEncryptPasswordException(context string) encryptPasswordException {
	return encryptPasswordException{
		SHttpException{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("encryptPasswordException fired in [%s]", context),
		},
	}
}

type databaseErrorException struct {
	SHttpException
}

func NewDatabaseErrorException(context string) databaseErrorException {
	return databaseErrorException{
		SHttpException{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("databaseErrorException fired in [%s]", context),
		},
	}
}
