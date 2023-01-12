package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/httputil"
	"github.com/yangwawa0323/gin-gorm-jwt/utils"
)

func MockJsonResponse(ctx *gin.Context, filename string) {
	jsonData, err := utils.ReadMockJson(filename)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": jsonData,
	})
}
