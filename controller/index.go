package controller

import (
	"net/http"

	"github.com/double1996/joplin-web-blog/models"
	"github.com/gin-gonic/gin"
)

//func Home(c *gin.Context) {
//	posts, err := models.ListAllPost("")
//	if err != nil {
//	}
//	c.HTML(http.StatusOK, "index/home.html", gin.H{
//		"posts": posts,
//		//"tags":            models.MustListTag(),
//		//"archives":        models.MustListPostArchives(),
//		//"links":           models.MustListLinks(),
//		//"user":            user,
//		//"pageIndex":       pageIndex,
//		//"totalPage":       int(math.Ceil(float64(total) / float64(pageSize))),
//		//"path":            c.Request.URL.Path,
//		//"maxReadPosts":    models.MustListMaxReadPost(),
//		//"maxCommentPosts": models.MustListMaxCommentPost(),
//	})
//}

func Resume(c *gin.Context) {
	resume, err := models.GetResumePost()
	if err != nil || resume == nil {
		Handle404(c)
		return
	}
	c.HTML(http.StatusOK, "index/resume.html", gin.H{
		"post": resume,
	})
}
