package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title          string    `json:"title"`
	PublishTime    time.Time `json:"publishTime" gorm:"column:publish_time"`
	Favorited      int32     `json:"favorited"`
	FavoritedUsers []*User   `json:"favorited_users" gorm:"many2many:user_favorited"`
}
