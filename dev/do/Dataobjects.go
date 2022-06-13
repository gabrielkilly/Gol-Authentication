package do

type HttpResponse[T any] struct {
	Code int
	Data T
}

type CreateAuthUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateAuthUserResponse struct {
	Id       string `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
