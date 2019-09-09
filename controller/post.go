package controller

import (
	"github.com/double1996/smart-evernote-blog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostsIndex(c *gin.Context) {
	posts, err := models.ListAllPost("")
	if err != nil {

	}
	c.HTML(http.StatusOK, "post/index", gin.H{
		"posts": posts, // TODO: need list comments
	})
}

func PostGetByID(c *gin.Context) {
	id := c.Param("id")
	post, err := models.GetPostByID(id)
	if err != nil {
	}
	post.View++

	tags, _ := models.GetPostByID(id)
	c.HTML(http.StatusOK, "", gin.H{
		"post": post,
		"tags": tags,
	})
}

func PostListByTag(c *gin.Context) {

}
