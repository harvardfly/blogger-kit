package models

type User struct {
	BaseModel
	Username string
	Password string
	Email    string
}

func (User) TableName() string {
	return "user"
}
