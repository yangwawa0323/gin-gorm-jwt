package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/auth"
	"github.com/yangwawa0323/gin-gorm-jwt/httputil"
	"github.com/yangwawa0323/gin-gorm-jwt/models"
	"github.com/yangwawa0323/gin-gorm-jwt/services"
)

// NewPage godoc
//
//		@Summary		Post a new Page
//		@Description	Create a new page and post to server
//		@Tags			page
//		@Accept			json
//		@Produce		json
//		@Param			page	body		models.Page	true	"Page with content"
//		@Success		200 {object}	models.ResponseMessage
//	 @Router			/api/page/new [post]
func NewPage(ctx *gin.Context) {
	var page models.Page
	// TODO: should get user info and post database
	// ctx.Header("Access-Control-Allow-Origin", "*")

	if err := ctx.ShouldBindJSON(&page); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	// Method 1: Globally use the database service
	// dbsvc := services.NewDBService()
	// record := dbsvc.DB.Create(&page)

	pgsvc := services.NewPageService(&page)
	// pgsvc.Page = &page
	// TODO:
	if err := pgsvc.New(&page); err != nil {
		if err != nil {
			httputil.NewError(ctx, http.StatusInternalServerError, err)
			return
		}
	}

	ctx.JSON(http.StatusOK, models.ResponseMessage{
		Code:    http.StatusOK,
		Message: "Save the page content to database successfully.",
	})
}

// Get all page and format a json file godoc
//
//	@Summary		List all pages
//	@Description	List all pages
//	@Tags			page
//	@Accept			json
//	@Produce		json
//	@Success		200 {array}		models.Page
//	@Failure		400	{object}	httputil.HTTPError
//	@Router			/api/page/all [get]
func AllPages(ctx *gin.Context) {
	var pages []models.Page

	// "&" in &page is very important, the Scan() must operate a section memory which
	// is assignable. so we provided the slice pointer.
	dbsvc := services.NewDBService()
	dbsvc.DB.Where("content is not null").Find(&pages)

	ctx.PureJSON(http.StatusOK, pages)
}

func NotImplemented(ctx *gin.Context) {
	ctx.Status(200)

	message := "<h1>Not implemented yet</h1>"

	claim, err := auth.ParseClaim(auth.ExtractTokenString(ctx))
	if err == nil {
		debug(err.Error())
		message += fmt.Sprintf("<h4>welcome %s</h4>", claim.Username)
	}
	message += fmt.Sprintf("<span>%s</span>", time.Now().Format("2006-01-02 15:04:05"))
	ctx.Writer.WriteString(message)
}
