package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/auth"
	"github.com/yangwawa0323/gin-gorm-jwt/models"
	"github.com/yangwawa0323/gin-gorm-jwt/services"
)

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

	var sendmail chan bool = make(chan bool)

	go func() {
		// TODO : send a activate mail
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
		sendmail <- mailDialer.SendMail_gomailV2() == nil

	}()

	hasBeenSent := <-sendmail
	if hasBeenSent {
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
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user login",
	})
}

func RefreshToken(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user refresh token",
	})
}

func ListMessages(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "list messages",
	})
}
