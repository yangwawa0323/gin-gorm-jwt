package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/models"
)

type TokenString struct {
	Token string `json:"token"`
}

// PingPong godoc
//
//		@Summary		Ping pong
//		@Description	Ping pong example, need the token authenticated
//		@Tags			authenticate
//		@Accept			json
//		@Produce		json
//		@Param			Authorization	header		string	true	"with the bearer stared"
//		@Success		200 {object}	models.ResponseMessage
//	 @Router			/api/secured/ping [post]
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.ResponseMessage{
		Code:    http.StatusOK,
		Message: "pong",
	})
}
