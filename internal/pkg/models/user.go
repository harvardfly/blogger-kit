package models

// User 用户model
type User struct {
	BaseModel
	Username string
	Password string
	Email    string
}

// TableName 获取表名
func (User) TableName() string {
	return "user"
}
