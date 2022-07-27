package routers

import (
	"go_reddit/controller"
	"go_reddit/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	v1 := r.Group("/api/v1")
	// 注册业务路由
	v1.POST("/signup", controller.SignUpController)
	// 登录业务路由
	v1.POST("/login", controller.SignInController)
	v1.GET("/refresh_token", controller.RefreshTokenController)

	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityController)

		v1.POST("/post", controller.PostController)

		v1.GET("/posts2", controller.PostListController)
		v1.GET("/post/:id", controller.PostDetailController)
		v1.POST("/vote", controller.VoteController)
	}

	r.GET("/hello", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		// 登录的用户才能访问hello， 判断请求头中是否有token
		c.JSON(http.StatusOK, gin.H{"message": "world"})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "404",
		})
	})
	return r
}
