package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	Poster  *User  `json:"poster"`
}

type QuestionReply struct {
	gorm.Model
	Question  *Question `json:"question"`
	ReplyUser *User     `json:"reply_user"`
	Content   string    `json:"content"`
}
