package main

import (
	"github.com/double1996/smart-evernote-blog/routers"
	"os"
)

var port = os.Getenv("PORT")

func main() {

	if port == "" {
		port = "2345"
	}

	routersInit := routers.InitRouter()

}
