package routes

import (
	"net/http"
	"web_app/controller"
	"web_app/logger"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	v1 := r.Group("/api/v1")

	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.GET("/refresh_token", controller.RefreshTokenHandler)

	v1.Use(controller.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/post2", controller.GetPost2ListHandler)
		v1.GET("/post", controller.GetPostListHandler)
		v1.GET("/post/")

		v1.POST("/vote", controller.VoteHandler)
		//v1.GET("/post2", controller.GetPostListHandler())
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
