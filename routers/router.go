package routers

import (
	"github.com/double1996/joplin-web-blog/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	posts := router.Group("/posts")
	{
		posts.GET("/", controller.PostsIndex)
		posts.GET("/:id", controller.PostGetContextByID)

	}
	router.GET("/", controller.PostsIndex)
	router.GET("/resume", controller.Resume)
	router.Static("/static", "./static")

	// login
	//router.GET("/signin", controllers.SigninGet)
	//router.POST("/signin", controllers.SigninPost)
	//router.GET("/logout", controllers.LogoutGet)
	//router.GET("/oauth2callback", controllers.Oauth2Callback)
	//router.GET("/auth/:authType", controllers.AuthGet)

	router.NoRoute(controller.Handle404)
}
