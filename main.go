package main

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/routers"
	"github.com/yangwawa0323/gin-gorm-jwt/services"
	"github.com/yangwawa0323/gin-gorm-jwt/utils"
	myerr "github.com/yangwawa0323/gin-gorm-jwt/utils/errors"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/yangwawa0323/gin-gorm-jwt/docs"
)

//go:embed assets/*.ico
var embeddedFiles embed.FS

// var errorDebug = utils.ErrorDebug
var debug = utils.Debug

// @title 51cloudclass Gin Swagger API
// @version 1.0.0
// @description This is a 51cloudclass.com web api server.
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {

	// Initialize Database
	dbsvc := services.NewDBService()
	dbsvc.Migrate()

	// utils.LoadDotEnv() // Load the dotenv configuration

	router := initRouter()

	done := make(chan os.Signal, 1)
	srv := gracefulHTTPServe(done, ":8080", router)
	<-done
	timeoutShutdown(srv)

}

func initRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8000", "http://localhost:3000", "http://google.com", "http://facebook.com"}
	config.AllowHeaders = []string{"*"}
	router.Use(cors.New(config))

	docs.SwaggerInfo.BasePath = "/"
	// by default `swag init` generate the file and put it to `./docs/swagger.json`
	// router.GET("/swagger.json", gin.WrapH(http.FileServer(http.Dir("./docs/"))))
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routers.RootRouter(router)

	// Favicon
	favFS := &routers.FavFS{
		FS:     &embeddedFiles,
		Engine: router,
	}
	routers.FavFSRouter(favFS)
	// Testing URL
	routers.TestUrl(router)
	// functional api
	routers.ApiRouter(router)

	return router
}

func gracefulHTTPServe(done chan os.Signal, port string, handler http.Handler) *http.Server {
	debug("Start server at", port)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    port,
		Handler: handler,
	}

	go func(tls bool) {

		if tls {
			cert, key := utils.GetCertFiles(utils.InitConfig())
			if err := srv.ListenAndServeTLS(cert, key); err != nil &&
				err != http.ErrServerClosed {

				// TODO: to some clean things.
				debug(myerr.Errors[myerr.ServicePortIsUsed], err.Error())
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				debug(myerr.Errors[myerr.ServicePortIsUsed], err.Error())
			}
		}
	}(false)

	return srv
}

func timeoutShutdown(srv *http.Server) {
	// debug("server stop...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		debug(fmt.Sprintf("server shutdown failed: %+v", err))
	}
	// debug("server exited properly")
}
