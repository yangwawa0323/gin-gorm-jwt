package models

import "gorm.io/gorm"

type Creation int64

const (
	Original Creation = iota
	Quote
)

type News struct {
	gorm.Model
	Title       string   `json:"title"`
	PublishTime string   `json:"publish_time"`
	Creation    Creation `json:"creation" gorm:"type:tinyint"`
	Content     string   `json:"content"`
}
