package routers

import (
	"github.com/double1996/joplin-web-blog/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	posts := router.Group("/posts")
	{
		posts.GET("/", controller.PostsIndex)
	}
	router.GET("/", controller.IndexGet)
	router.GET("/about", controller.About)
	router.Static("/static", "./static")

	router.NoRoute(controller.Handle404)
}
