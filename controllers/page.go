package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewPage(ctx *gin.Context) {
	// TODO: should get user info and post data
	//ctx.ShouldBindJSON()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Save the page content to database [fake].",
	})
}
