package mocks

import (
	"gol/the-basics/dev/do"

	"github.com/gin-gonic/gin"
)

var TimesCalled int
var LastError error
var LastResponseCode int

type responseMapperMock struct {
	TimesCalled      int
	LastError        error
	LastResponseCode int
}

var ResponseMapperMock *responseMapperMock

func MapResponseMock[T any](httpResponse do.HttpResponse[T], err error) func(*gin.Context) {
	(*ResponseMapperMock).TimesCalled++
	(*ResponseMapperMock).LastError = err
	(*ResponseMapperMock).LastResponseCode = httpResponse.Code

	return func(ctx *gin.Context) {}
}

func (this *responseMapperMock) clear() {
	ResponseMapperMock = &responseMapperMock{}
}
