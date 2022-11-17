package models

import "gorm.io/gorm"

type Page struct {
	gorm.Model
	Content string `json:"content" binding:"required" gorm:"content"`
}
