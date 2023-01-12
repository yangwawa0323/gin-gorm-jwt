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
		// api/user/activate-by-email/%d?token=
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.POST("/ping", controllers.Ping)
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
		page.PUT("/update/:page_id", controllers.NotImplemented)
		page.DELETE("/delete/:page_id", controllers.NotImplemented)

		page.GET("/add-favorite/:page_id", controllers.NotImplemented)
		page.GET("/revoke-favorite/:page_id", controllers.NotImplemented)

		page.GET("/relative-pages", controllers.NotImplemented)

		page.POST("/add-comment/:pageID", controllers.NotImplemented)

	}

}

func UserRouter(base *gin.RouterGroup) {
	user := base.Group("/user", middlewares.Auth())
	{
		// api/user/activate-by-email/%d?token=
		user.GET("/activate-by-email/:user_id", controllers.ConfirmMailActivate)
		user.PUT("/disable/:user_id")

		user.GET("/profile/:name/:user_id", controllers.GetUserProfile)

		user.POST("/change-password", controllers.ChangePassword)
		user.GET("/list/messages", controllers.ListMessages)
		user.POST("/upload-avatar", controllers.UploadAvator)
	}
}

func CourseRouter(base *gin.RouterGroup) {
	course := base.Group("/course")
	{
		course.POST("/new", controllers.NotImplemented)
		course.GET("/lists", controllers.NotImplemented)
		course.PUT("/update/:course_id", controllers.NotImplemented)
		course.DELETE("/delete/:course_id", controllers.NotImplemented)

		course.GET("/search/:keyword", controllers.NotImplemented)
	}
}
