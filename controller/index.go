package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexGet(c *gin.Context) {
	//post := models.GetPostByID()
	c.HTML(http.StatusOK, "index/index.html", gin.H{
		//"posts":           posts,
		//"tags":            models.MustListTag(),
		//"archives":        models.MustListPostArchives(),
		//"links":           models.MustListLinks(),
		//"user":            user,
		//"pageIndex":       pageIndex,
		//"totalPage":       int(math.Ceil(float64(total) / float64(pageSize))),
		//"path":            c.Request.URL.Path,
		//"maxReadPosts":    models.MustListMaxReadPost(),
		//"maxCommentPosts": models.MustListMaxCommentPost(),
	})
}

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "index/about.html", gin.H{})
}
