package do

import "net/http"

type HttpResponse[T any] struct {
	Code int
	Data T
}

func HttpOk[T any](data T) HttpResponse[T] {
	return HttpResponse[T]{
		Code: http.StatusOK,
		Data: data,
	}
}

func HttpCreated[T any](data T) HttpResponse[T] {
	return HttpResponse[T]{
		Code: http.StatusCreated,
		Data: data,
	}
}

func Empty[T any]() HttpResponse[T] {
	return HttpResponse[T]{}
}
