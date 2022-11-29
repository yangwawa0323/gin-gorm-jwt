package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/models"
	"github.com/yangwawa0323/gin-gorm-jwt/routers"
	"github.com/yangwawa0323/gin-gorm-jwt/services"
)

func main() {
	// Initialize Database
	dbsvc := services.NewDBService()
	dbsvc.Migrate()

	router := initRouter()

	models.SendMail_gomailV2()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8000", "http://localhost:3000", "http://google.com", "http://facebook.com"}
	// config.AllowAllOrigins = true
	config.AllowHeaders = []string{"*"}
	// cors.New(config)
	router.Use(cors.New(config))
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()

	// Testing URL
	routers.TestUrl(router)

	// functional api
	routers.ApiRouter(router)

	return router
}
