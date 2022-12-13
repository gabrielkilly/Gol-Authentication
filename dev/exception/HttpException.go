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
}

func (this SHttpException) Error() string {
	return fmt.Sprintf("HttpException status %d with message: %s", this.Code, this.Message)
}
