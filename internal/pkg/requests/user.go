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

type FindByIDRequest struct {
	ID int
}

type FindByEmailRequest struct {
	Email string
}
