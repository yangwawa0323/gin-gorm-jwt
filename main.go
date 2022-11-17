package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/database"
	"github.com/yangwawa0323/gin-gorm-jwt/routers"
	"github.com/yangwawa0323/gin-gorm-jwt/utils"
)

func main() {
	// Initialize Database
	database.Connect(utils.GetConnectURI())
	database.Migrate()

	router := initRouter()

	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000", "http://127.0.0.1:3000"}
	config.AllowOrigins = []string{"http://localhost:3000", "http://google.com", "http://facebook.com"}
	// config.AllowAllOrigins = true
	config.AllowHeaders = []string{"*"}
	// cors.New(config)
	router.Use(cors.New(config))
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	routers.ApiRouter(router)

	return router
}
