package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/auth"
	"github.com/yangwawa0323/gin-gorm-jwt/models"
	"github.com/yangwawa0323/gin-gorm-jwt/services"
	"github.com/yangwawa0323/gin-gorm-jwt/utils"
	"gorm.io/gorm"
)

var Errors = utils.Errors

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(ctx *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": Errors[utils.BindPostDataError],
		})
		ctx.Abort()
		return
	}

	dbsvc := services.NewDBService()

	record := dbsvc.DB.Where("email = ?", request.Email).First(&user)
	if record.Error != nil || record.Error == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": Errors[utils.ErrRecordNotFound],
		})
		ctx.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": Errors[utils.CredentialError],
		})
		ctx.Abort()
		return
	}

	// debug(fmt.Sprintf("email : %s, password: %s \n", user.Email, user.Password))

	tokenString, err := auth.GenerateToken(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": Errors[utils.GenerateTokenError],
		})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}
