package models

import (
	"time"
)

// BaseModel .
type BaseModel struct {
	ID        int       `gorm:"primary_key" json:"id" form:"id"`
	CreatedAt time.Time `form:"created_at" json:"created_at"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at"`
}
