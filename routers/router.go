package routers

import (
	"github.com/double1996/smart-evernote-blog/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	posts := router.Group("/posts")
	{
		posts.GET("/", controller.PostsIndex)
	}
	en := router.Group("/evernote") // 印象笔记路由组
	{
		en.GET("/webhook/", controller.GetNewEverNoteByWebhook)
		en.GET("/callback", controller.GetEverNoteCallback)
	}
	router.GET("/", controller.IndexGet)
	router.GET("/about", controller.About)
	router.Static("/static", "./static")

	router.NoRoute(controller.Handle404)
}
