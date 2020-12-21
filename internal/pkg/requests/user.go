package requests

// LoginRequest 登录请求参数
type LoginRequest struct {
	Email    string
	Password string
}

// RegisterRequest 注册请求参数
type RegisterRequest struct {
	Username string
	Email    string
	Password string
}

// FindByIDRequest 通过ID查找 请求参数
type FindByIDRequest struct {
	ID int
}

// FindByEmailRequest 通过Email查找 请求参数
type FindByEmailRequest struct {
	Email string
}
