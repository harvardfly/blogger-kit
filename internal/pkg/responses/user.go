package responses

// UserInfo 用户信息返回值
type UserInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// RegisterUser 注册用户返回值
type RegisterUser struct {
	Username string
	Password string
	Email    string
}

// LoginResponse 定义登录返回结构体
type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	ExpireAt    int64  `json:"expireAt"`
	TimeStamp   int64  `json:"timeStamp"`
}

// RegisterResponse 注册返回值
type RegisterResponse struct {
	UserInfo *UserInfo `json:"user_info"`
}
