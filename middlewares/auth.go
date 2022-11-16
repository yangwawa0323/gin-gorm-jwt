package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/auth"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		log.Printf("[DEBUG]: tokenString %s\n", tokenString)
		if tokenString == "" {
			ctx.JSON(401, gin.H{
				"error": "request does not contain an access token"},
			)
			ctx.Abort()
			return
		}
		err := auth.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(401, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
