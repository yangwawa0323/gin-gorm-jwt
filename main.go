package main

import (
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
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	routers.ApiRouter(router)

	return router
}
