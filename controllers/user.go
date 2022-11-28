package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	usersvc := services.NewUserService(&user)
	// usersvc.User = &user
	record := usersvc.Save()
	if record != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": record.Error(),
		})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"userId":   user.ID,
		"email":    user.Email,
		"username": user.Username,
	})
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
