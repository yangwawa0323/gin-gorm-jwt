package models

import (
	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Gender int64

const (
	Famale Gender = iota
	Male
	Unclosed
)

type Privilege int64

const (
	Readable Privilege = 1 << iota
	Writeable
	Downloadable
	Uploadable
)

type UserClass int64

const (
	Guest UserClass = iota
	MonthlySubscription
	AnnualSubscription
)

type User struct {
	gorm.Model
	Name              string     `json:"name"`
	Username          string     `json:"username" gorm:"unique"`
	Email             string     `json:"email" gorm:"unique"`
	Password          string     `json:"password"`
	Phone             string     `json:"phone"`
	IdentityNumber    string     `json:"identity_number"`
	Privilege         Privilege  `json:"privilege" gorm:"type:tinyint"`
	Gender            Gender     `json:"gender" gorm:"type:tinyint"`
	UserClass         UserClass  `json:"user_class" gorm:"type:tinyint"`
	FavoritedArticles []*Article `gorm:"many2many:user_favorited"`
	Followee          []*User    `gorm:"many2many:user_followee"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	return err
}

func (user *User) RefreshToken(ctx *gin.Context) (token []byte, err error) {
	token = []byte("")
	err = errors.New("not implemented yet")
	return
}
