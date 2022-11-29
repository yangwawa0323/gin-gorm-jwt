package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/models"
	"github.com/yangwawa0323/gin-gorm-jwt/services"
)

func NewPage(ctx *gin.Context) {
	var page *models.Page
	// TODO: should get user info and post database
	// ctx.Header("Access-Control-Allow-Origin", "*")

	if err := ctx.ShouldBindJSON(page); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	// Method 1: Globally use the database service
	// dbsvc := services.NewDBService()
	// record := dbsvc.DB.Create(&page)

	pgsvc := services.NewPageService(page)
	// pgsvc.Page = &page
	// TODO:
	if err := pgsvc.New(page); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Save the page content to database successfully.",
	})
}

// Get all page and format a json file
func AllPages(ctx *gin.Context) {
	var pages []models.Page

	// "&" in &page is very important, the Scan() must operate a section memory which
	// is assignable. so we provided the slice pointer.
	dbsvc := services.NewDBService()
	dbsvc.DB.Where("content is not null").Find(&pages)

	ctx.PureJSON(http.StatusOK, pages)
}
