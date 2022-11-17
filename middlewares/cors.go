package middlewares

import (
	"os"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOW_ORIGIN"))
		ctx.Next()
	}
}
