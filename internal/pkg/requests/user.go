package requests

type LoginRequest struct {
	Email    string
	Password string
}

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}
