package responses

import (
	"time"
)

// Article response 数据结构
type Article struct {
	ID    int    `json:"id"`
	Title string `json:"Title"`
}

// ArticleRes response 数据结构
type ArticleRes struct {
	ID        int       `json:"id"`
	Summary   string    `json:"summary"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Category  Category  `json:"category"`
}
