package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestUrl(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome")
	})
	router.GET("/hello", Helloworld)
	router.GET("health", HealthCheck)
}

// A get function which returns a hello world string by json
// @BasePath /

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "helloworld")
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func HealthCheck(ctx *gin.Context) {
	res := map[string]any{
		"data": "Service is up and running",
	}

	ctx.JSON(http.StatusOK, res)
}
