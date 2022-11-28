package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestUrl(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome")
	})
}
