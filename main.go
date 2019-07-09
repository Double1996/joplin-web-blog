package main

import (
	"github.com/double1996/smart-evernote-blog/config"
	"os"

	"github.com/gin-gonic/gin"
)

var port = os.Getenv("PORT")

func main() {
	engine := gin.New()
	engine.Use(gin.Logger())
	loadTemplate(engine)
}

func loadTemplate(e *gin.Engine) {
	root := config.Root
	e.StaticFS("/public")
	e.LoadHTMLGlob(root + "/templates/*")
}
