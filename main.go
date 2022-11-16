package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/controllers"
	"github.com/yangwawa0323/gin-gorm-jwt/database"
	"github.com/yangwawa0323/gin-gorm-jwt/middlewares"
	"github.com/yangwawa0323/gin-gorm-jwt/utils"
)

func main() {
	// Initialize Database
	database.Connect(utils.GetConnectURI())
	database.Migrate()

	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
