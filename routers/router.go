package routers

import (
	"github.com/double1996/smart-evernote-blog/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	posts := router.Group("/posts")
	{
		posts.GET("/", controller.GetPost)
	}
}
