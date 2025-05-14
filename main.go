package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// db variable creation

func main() {
	router := gin.Default()
	// Set up session store
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.html")
	//login page
	router.GET("", loginpage)
	router.GET("/index", index)
	router.GET("/firstpage", firstpage)
	router.Run(":8080")

}

// AuthMiddleware validates session and user authentication
func loginpage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
