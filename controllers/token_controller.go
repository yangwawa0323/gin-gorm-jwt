package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/auth"
	"github.com/yangwawa0323/gin-gorm-jwt/httputil"
	"github.com/yangwawa0323/gin-gorm-jwt/models"
	"github.com/yangwawa0323/gin-gorm-jwt/services"
	myerr "github.com/yangwawa0323/gin-gorm-jwt/utils/errors"
	"gorm.io/gorm"
)

var Errors = myerr.Errors

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	models.ResponseMessage
	Token string `json:"token"`
}

// GenerateToken godoc
//
//		@Summary		Generate Token
//		@Description	Generate Token for the authenticated user.
//		@Tags			authenticate
//		@Accept			json
//		@Produce		json
//		@Param			request	body		TokenRequest	true	"email and password"
//		@Success		200 {object}	TokenResponse
//	 @Router			/api/token [post]
func GenerateToken(ctx *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := ctx.ShouldBindJSON(&request); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New(Errors[myerr.BindPostDataError]))
		return
	}

	dbsvc := services.NewDBService()

	record := dbsvc.DB.Where("email = ?", request.Email).First(&user)
	if record.Error != nil || record.Error == gorm.ErrRecordNotFound {
		httputil.NewError(ctx, http.StatusInternalServerError, errors.New(Errors[myerr.ErrRecordNotFound]))
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		httputil.NewError(ctx, http.StatusUnauthorized, errors.New(Errors[myerr.CredentialError]))
		return
	}

	// debug(fmt.Sprintf("email : %s, password: %s \n", user.Email, user.Password))

	tokenString, err := auth.GenerateToken(&user)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, errors.New(Errors[myerr.GenerateTokenError]))
		return
	}
	ctx.JSON(http.StatusOK, TokenResponse{
		ResponseMessage: models.ResponseMessage{
			Code:    http.StatusOK,
			Message: "Generate token successfully.",
		},
		Token: tokenString,
	})

}
