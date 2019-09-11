package models

import "time"

type Chat struct {
	Username  string    `json:"username"`
	Message   string    `json:"message"`
	CratedAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MessageList struct {
	Messages []string `json:"messages"`
}
