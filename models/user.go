package models

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const mailTemplateFile = "./templates/activate-mail.html"

var errorDebug = utils.ErrorDebug

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
	Deleteable
	Downloadable
	Uploadable
	ChangeUserClass
	Grant
	Admin Privilege = Readable | Writeable | Deleteable |
		Downloadable | Uploadable | ChangeUserClass | Grant
)

var LiteralPrivilege = map[int64]string{
	int64(Readable):        "readable",
	int64(Writeable):       "writeable",
	int64(Deleteable):      "deleteable",
	int64(Downloadable):    "downloadable",
	int64(Uploadable):      "uploadable",
	int64(ChangeUserClass): "change user class",
	int64(Grant):           "grant privilege",
	int64(Admin):           "administration",
}

type UserClass int64

const (
	Guest UserClass = iota
	MonthlySubscription
	AnnualSubscription
	Administrator
)

var LiteralUserClass = map[int64]string{
	int64(Guest):               "guest",
	int64(MonthlySubscription): "monthly subscription user",
	int64(AnnualSubscription):  "annual subscription user",
	int64(Administrator):       "Administrator",
}

type User struct {
	gorm.Model
	Name              string     `json:"name" gorm:"column:name;type:varchar(100)"`
	Username          string     `json:"username" gorm:"not null;type:varchar(100)"`
	Email             string     `json:"email" gorm:"unique"`
	Password          string     `json:"password" gorm:"type:varchar(255)"`
	AvatarURL         string     `json:"avatar_url" gorm:"type:varchar(255)" example:"https://picsum.photos/200/300?random=1"`
	IdentityNumber    string     `json:"identity_number" gorm:"type:char;size:18"`
	Privilege         Privilege  `json:"privilege" gorm:"type:tinyint"`
	Gender            Gender     `json:"gender" gorm:"type:tinyint"`
	UserClass         UserClass  `json:"user_class" gorm:"type:tinyint"`
	FavoritedArticles []*Article `json:"favorited_articles" gorm:"many2many:user_favorited"`
	Followee          []*User    `json:"followee" gorm:"many2many:user_followee"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
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

// Should move to controller/user.go
func (user *User) RefreshToken(ctx *gin.Context) (token []byte, err error) {
	token = []byte("")
	err = errors.New("not implemented yet")
	return
}

func (user *User) Secret() (secret []byte) {
	secret = []byte(fmt.Sprintf("%s@51cloudclass@%s", user.Email, user.Password))
	return
}

func (user *User) GenerateActivateMailBody() (string, error) {

	activateString, err := bcrypt.GenerateFromPassword(
		user.Secret(),
		bcrypt.DefaultCost)

	if err != nil {
		return "", errors.New("error to generate activate mail string")
	}

	var buf *bytes.Buffer = new(bytes.Buffer)
	// TODO: hard code here
	srvHost := os.Getenv("SERVER_URL")
	url :=
		fmt.Sprintf("%s/api/user/activate-by-email/%d?token=%s",
			srvHost, user.ID,
			url.QueryEscape(string(activateString)),
		)

	tmpl, err := template.ParseFiles(mailTemplateFile)
	if err != nil {
		errorDebug(err, "\n[DEBUG]can not parse the template file")
		return "", err
	}

	if err := tmpl.Execute(buf, url); err != nil {
		errorDebug(err,
			"\n[DEBUG]can not execute applies a parsed template to specified data object\n\n")
		return "", err
	}
	return buf.String(), nil
}

func (user *User) IsActivateMailStringValid(activateString []byte) bool {
	err := bcrypt.CompareHashAndPassword(activateString, user.Secret())
	return err == nil
}
