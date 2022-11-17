package middlewares

import (
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Next()
	}
}
