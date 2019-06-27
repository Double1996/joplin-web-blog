package main

import (
	"os"
	"http"
)

var port = os.Getenv("PORT")

func main() {
	
	if port == "" {
		port = "2345"
	}
	mux := http.Newser
}

type Site struct {
	Post modles.
}




