package models

type Article struct {
	Base
	Title    string `json:"title"`
	Username string `json:"username"`
	Content  string `json:"content"`
	UserID   int    `json:"user_id"`
}
