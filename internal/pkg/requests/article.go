package requests

import (
	"time"

	"pkg.zpf.com/golang/blogger-kit/internal/pkg/responses"
)

// Article request 数据结构
type Article struct {
	CategoryID int    `form:"category_id" json:"category_id" validate:"required"`
	Summary    string `form:"summary" json:"summary" validate:"required"`
	Title      string `form:"title" json:"title" validate:"required"`
	UserName   string `form:"username" json:"username"`
}

// ArticleInfo request 数据结构
type ArticleInfo struct {
	ID int `form:"id" json:"id" validate:"required"`
}

// ArticleEdit request 数据结构
type ArticleEdit struct {
	ID         int    `form:"id" json:"id" validate:"required"`
	CategoryID int    `form:"category_id" json:"category_id" validate:"required"`
	Summary    string `json:"summary" validate:"required"`
	Title      string `json:"title" validate:"required"`
}

// ArticleCategoryEdit request 数据结构
type ArticleCategoryEdit struct {
	ID         int `form:"id" json:"id" validate:"required"`
	CategoryID int `form:"category_id" json:"category_id" validate:"required"`
}

// ArticleES request 数据结构
type ArticleES struct {
	ID        int                `json:"id" validate:"required"`
	Summary   string             `json:"summary" validate:"required"`
	Title     string             `json:"title" validate:"required"`
	Category  responses.Category `json:"category" validate:"required"`
	CreatedAt time.Time          `json:"created_at" validate:"required"`
	UpdatedAt time.Time          `json:"updated_at" validate:"required"`
}
