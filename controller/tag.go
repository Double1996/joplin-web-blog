package controller

//func TagGet(c *gin.Context) {
//	var (
//		tagName   string
//		page      string
//		pageIndex int
//		pageSize  = system.GetConfiguration().PageSize
//		total     int
//		err       error
//		posts     []*models.Post
//	)
//	tagName = c.Param("tag")
//	page = c.Query("page")
//	pageIndex, _ = strconv.Atoi(page)
//	if pageIndex <= 0 {
//		pageIndex = 1
//	}
//	posts, err = models.ListPublishedPost(tagName, pageIndex, pageSize)
//	if err != nil {
//		c.AbortWithStatus(http.StatusInternalServerError)
//		return
//	}
//	total, err = models.CountPostByTag(tagName)
//	if err != nil {
//		c.AbortWithStatus(http.StatusInternalServerError)
//		return
//	}
//	policy = bluemonday.StrictPolicy()
//	for _, post := range posts {
//		post.Tags, _ = models.ListTagByPostId(strconv.FormatUint(uint64(post.ID), 10))
//		post.Body = policy.Sanitize(string(blackfriday.MarkdownCommon([]byte(post.Body))))
//	}
//	c.HTML(http.StatusOK, "index/index.html", gin.H{
//		"posts":           posts,
//		"tags":            models.MustListTag(),
//		"archives":        models.MustListPostArchives(),
//		"links":           models.MustListLinks(),
//		"pageIndex":       pageIndex,
//		"totalPage":       int(math.Ceil(float64(total) / float64(pageSize))),
//		"maxReadPosts":    models.MustListMaxReadPost(),
//		"maxCommentPosts": models.MustListMaxCommentPost(),
//	})
//}
