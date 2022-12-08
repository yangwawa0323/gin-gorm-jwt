package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/controllers"
)

func RootRouter(router *gin.Engine) {

	// There two routes no need Auth middleware.
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
}
