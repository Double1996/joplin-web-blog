package controller

import (
	"github.com/double1996/joplin-web-blog/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PostsIndex(c *gin.Context) {
	var (
		pageIndex int
		pageSize  int
		page      string
	)
	page = c.Query("page")
	pageIndex, _ = strconv.Atoi(page)
	posts, err := models.ListAllPost("", pageIndex, pageSize)
	if err != nil {

	}
	c.HTML(http.StatusOK, "post/list.html", gin.H{
		"posts": posts, // TODO: need list comments
	})
}

func PostGetContextByID(c *gin.Context) {
	id := c.Param("id")
	post, err := models.GetPostByID(id)
	if err != nil {
	}
	post.View++
	err = post.UpdateView()
	if err != nil {

	}
	//tags, _ := models.GetPostByID(id)
	c.HTML(http.StatusOK, "post/display.html", gin.H{
		"post": post,
		//"tags": tags,
	})
}

func PostListByTag(c *gin.Context) {
	//tag := c.Param("tag")
}
