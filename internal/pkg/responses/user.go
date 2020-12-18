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

// LoginResponse 定义登录返回结构体
type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	ExpireAt    int64  `json:"expireAt"`
	TimeStamp   int64  `json:"timeStamp"`
}

type RegisterResponse struct {
	UserInfo *UserInfo `json:"user_info"`
}
