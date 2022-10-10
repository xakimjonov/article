package modules


import "time"



type Content struct {
	Title string `json:"title"   binding:"required"`
	Body  string `json:"body"   binding:"required"`
}

type Article struct {
	Id     string `json:"id" `
	AuthorId string `json:"Author_id" binding:"required" `
	Content
	CreatedAt time.Time  `json:"created_at" `
	UpdatedAt *time.Time `json:"updated_at" `
}

