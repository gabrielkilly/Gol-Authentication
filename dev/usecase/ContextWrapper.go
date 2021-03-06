package usecase

import (
	"errors"
	"fmt"
	"gol/the-basics/dev/do"
	"gol/the-basics/dev/exception"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServiceContext[T any](response func(*gin.Context) (do.HttpResponse[T], error)) func(*gin.Context) {
	httpResponse, err := response()
	if err != nil {
		var httpError exception.SHttpException
		if errors.As(err, &httpError) {
			return func(ctx *gin.Context) {
				sendHttpErrorMessage(ctx, httpError)
			}
		} else {
			return func(ctx *gin.Context) {
				sendErrorMessage(ctx, fmt.Errorf("Unrecognized error: %w", err))
			}
		}
	} else {
		return func(ctx *gin.Context) {
			sendData(ctx, httpResponse)
		}
	}
}

func sendData[T any](ctx *gin.Context, response do.HttpResponse[T]) {
	ctx.JSON(
		response.Code,
		response.Data,
	)
}

func sendErrorMessage(ctx *gin.Context, e error) {
	log.Printf("Error: %s", e.Error())
	ctx.JSON(
		http.StatusInternalServerError,
		gin.H{"message": "There is a problem with you request"},
	)
}

func sendHttpErrorMessage(ctx *gin.Context, e exception.SHttpException) {
	log.Printf("Error: %s", e.Error())
	ctx.JSON(
		e.Code,
		gin.H{"message": "There is a problem with you request"},
	)
}
