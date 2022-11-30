package controllers

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/auth"
	"github.com/yangwawa0323/gin-gorm-jwt/models"
	"github.com/yangwawa0323/gin-gorm-jwt/services"
	"github.com/yangwawa0323/gin-gorm-jwt/utils"
	"gorm.io/gorm"
)

var debug = utils.Debug
var errorDebug = utils.ErrorDebug

func RegisterUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	var saved chan bool = make(chan bool)
	var sent chan bool = make(chan bool)

	go func() {
		// TODO : send a activate mail
		<-saved
		body, err := user.GenerateActivateMailBody()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		mailDialer := models.NewMailDialer(
			"Welcome to register 51cloudclass website",
			body,
			user,
		)
		sent <- mailDialer.SendMail_gomailV2() == nil

	}()

	go func() {

		usersvc := services.NewUserService(&user)
		user.UserClass = models.Guest // new user by default is a **guest**
		result := usersvc.New(&user)
		if result != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error(),
			})
			ctx.Abort()
			return
		}

		saved <- result == nil
	}()

	if <-sent {
		token, err := auth.GenerateJWT(user.Email, user.Username)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"userId":   user.ID,
			"email":    user.Email,
			"username": user.Username,
			"token":    token,
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "fail to send user the activate mail",
		})
	}
}

func ConfirmMailActivate(ctx *gin.Context) {
	type Secret struct {
		Token string `form:"token"`
	}
	var svcUser = models.User{}
	var mailSecret Secret
	userIDparam, ok := ctx.Params.Get("user_id")

	if err := ctx.ShouldBindQuery(&mailSecret); err != nil {
		errorDebug(err)
	}
	userID, err := strconv.Atoi(userIDparam)

	if ok && err == nil {
		usersvc := services.NewUserService(&svcUser)
		user, dbErr := usersvc.FindUserByID(int64(userID))
		unescapedSecret, tokenErr := url.QueryUnescape(mailSecret.Token)

		if dbErr == nil && tokenErr == nil &&
			user.IsActivateMailStringValid([]byte(unescapedSecret)) {
			ctx.Redirect(http.StatusTemporaryRedirect, "/")
		} else {
			debug("[DEBUG]: activate token is invalid")
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid token to activate your account or has expired",
			})
		}
	}
}

func UploadAvator(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "upload avator",
	})
}

func ChangePassword(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "change passwod",
	})
}

func Login(ctx *gin.Context) {
	var user models.User = models.User{}
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, &user)
	} else {
		usersvc := services.NewUserService(&user)
		result := usersvc.DB.Where("name = ? and password = ?",
			usersvc.User.Name,
			usersvc.User.Password).First(&user)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user login",
	})
}

func ListMessages(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "list messages",
	})
}
