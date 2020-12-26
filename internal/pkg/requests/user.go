package requests

// LoginRequest 登录请求参数
type LoginRequest struct {
	Email    string `form:"email" json:"email" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

// RegisterRequest 注册请求参数
type RegisterRequest struct {
	Username string `form:"username" json:"username" validate:"required"`
	Email    string `form:"email" json:"email" validate:"required,email"`
	Password string `form:"password" json:"password" validate:"required"`
}

// FindByIDRequest 通过ID查找 请求参数
type FindByIDRequest struct {
	ID int `form:"id" json:"id" validate:"required"`
}

// FindByEmailRequest 通过Email查找 请求参数
type FindByEmailRequest struct {
	Email string `form:"email" json:"email" validate:"required"`
}
