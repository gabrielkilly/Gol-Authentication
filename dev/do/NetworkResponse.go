package do

type HttpResponse[T any] struct {
	Code int
	Data T
}

func EmptyResponse[T any]() HttpResponse[T] {
	return HttpResponse[T]{}
}
