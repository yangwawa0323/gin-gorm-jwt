package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/database"
	"github.com/yangwawa0323/gin-gorm-jwt/models"
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

	record := database.Instance.Create(&user)
	if record.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": record.Error.Error(),
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
