package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handle404(c *gin.Context) {
	HandleMessage(c, "Sorry, not found page!")
}

func HandleMessage(c *gin.Context, msg string) {
	c.HTML(http.StatusNotFound, "errors/error.html", gin.H{
		"message": msg,
	})
}

func writeJSON(ctx *gin.Context, h gin.H) {
	if _, ok := h["succeed"]; !ok {
		h["succeed"] = false
	}
	ctx.JSON(http.StatusOK, h)
}


