package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/auth"
	"github.com/yangwawa0323/gin-gorm-jwt/models"
	"github.com/yangwawa0323/gin-gorm-jwt/services"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(ctx *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	dbsvc := services.NewDBService()

	record := dbsvc.DB.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": record.Error.Error()},
		)
		ctx.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		ctx.Abort()
		return
	}

	// debug(fmt.Sprintf("email : %s, password: %s \n", user.Email, user.Password))

	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}

func VerifyToken(ctx *gin.Context) {

}

func RefreshToken(ctx *gin.Context) {

}
