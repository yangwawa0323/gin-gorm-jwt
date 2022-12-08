package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/auth"
	"github.com/yangwawa0323/gin-gorm-jwt/utils"
	myerr "github.com/yangwawa0323/gin-gorm-jwt/utils/errors"
)

var debug = utils.Debug
var Errors = utils.Errors

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		signedString := auth.ExtractTokenString(ctx)
		if _, err := auth.ParseClaim(signedString); err != nil {
			ctx.JSON(401, gin.H{
				"error": Errors[myerr.TokenIsInvalid],
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
