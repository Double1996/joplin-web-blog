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
	router.GET("/", controller.Home)
	router.GET("/resume", controller.Resume)
	router.Static("/static", "./static")

	router.NoRoute(controller.Handle404)
}
