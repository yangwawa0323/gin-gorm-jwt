package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/controllers"
	"github.com/yangwawa0323/gin-gorm-jwt/middlewares"
)

// ApiRouter extends the gin.Engine which gin.IRouterGroup
func ApiRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		// api.POST("/user/register", controllers.RegisterUser)
		// api/user/%d/activate-by-email?token=
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
		UserRouter(api)
		PageRouter(api)
	}

}

func PageRouter(base *gin.RouterGroup) {
	page := base.Group("/page")
	{
		page.GET("/view", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "test view page",
			})
		})

		// IMPORTANT: CORS() should before the RateLimit()
		page.POST("/new", middlewares.CORS(),
			middlewares.RateLimit(middlewares.POST_NEW_PAGE, 30*time.Second),
			controllers.NewPage)

		page.GET("/all", middlewares.CORS(), controllers.AllPages)

	}

}

func UserRouter(base *gin.RouterGroup) {
	user := base.Group("/user")
	{
		user.POST("/register", controllers.RegisterUser)
		// api/user/%d/activate-by-email?token=
		user.GET("/:user_id/activate-by-email", controllers.ConfirmMailActivate)
		user.POST("/login", controllers.Login)
		user.POST("/change-password", controllers.ChangePassword)
		user.GET("/list/messages", controllers.ListMessages)
		user.GET("/upload-avatar", controllers.UploadAvator)
	}
}
