package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/models"
)

func NewPage(ctx *gin.Context) {
	var page models.Page
	// TODO: should get user info and post database

	if err := ctx.ShouldBindJSON(&page); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	log.Printf("[DEBUG]: page: %#v\n", page)
	// record := database.Instance.Create(&page)

	// if record.Error != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": record.Error.Error(),
	// 	})
	// 	ctx.Abort()
	// 	return
	// }

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Save the page content to database successfully.",
	})
}
