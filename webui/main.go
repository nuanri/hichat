package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/signup/requests", func(c *gin.Context) {
		c.HTML(http.StatusOK, "auth_signup.html", gin.H{})
	})

	r.GET("/auth/signin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "auth_signin.html", gin.H{})
	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.Run(":8888")
}
