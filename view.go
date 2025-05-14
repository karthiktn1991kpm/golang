package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func firstpage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
