package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

var port = os.Getenv("PORT")

func main() {

	if port == "" {
		port = "2345"
	}

	app := gin.New()
	app.Use(gin.Logger())

}
