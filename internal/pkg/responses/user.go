package responses

type UserInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type RegisterUser struct {
	Username string
	Password string
	Email    string
}

type LoginResponse struct {
	UserInfo *UserInfo `json:"user_info"`
}

type RegisterResponse struct {
	UserInfo *UserInfo `json:"user_info"`
}
