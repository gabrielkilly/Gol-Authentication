package exception

import (
	"fmt"
)

type SHttpException struct {
	Code    int
	Message string
}

type IHttpException interface {
	Error() string
	StatusCode() int
}

func (exception SHttpException) Error() string {
	return fmt.Sprintf("HttpException status %d with message: %s", exception.Code, exception.Message)
}

func (exception SHttpException) StatusCode() int {
	return exception.Code
}
