package models

import (
	"time"

	"gorm.io/gorm"
)

type AuditLog struct {
	gorm.Model
	Timestamp time.Time `json:"timestamp"`
	Content   string    `json:"content"`
}
