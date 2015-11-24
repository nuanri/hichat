package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"nuanri/hichat/webui/apps/auth"
	"nuanri/hichat/webui/apps/message"
	"nuanri/hichat/webui/apps/middleware"
	"nuanri/hichat/webui/utils"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		session, err := middleware.GetSession(c)
		if err != nil {
			c.String(400, err.Error())
			c.Abort()
			return
		}

		sid := session.Sid
		if sid == "" {
			c.Redirect(302, "/auth/signin")
		}

		fmt.Println("sid===>", sid)
		//		if sid := session.Sid; sid = "" {
		//			c.Redirect(302, "/auth/signin")
		//		}
		// Set example variable
		if session != nil {
			c.Set("Sid", session.Sid)
			c.Set("User", session.User)
		}
		// before request
		c.Next()
		// after request
	}
}

func main() {
	utils.InitDB()

	r := gin.Default()

	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/signup/request", auth.SignupRequest)
	r.GET("/auth/signin", auth.GetSignin)
	r.POST("/auth/signin", auth.PostSignin)
	r.GET("/auth/signout", auth.Siginout)

	auth := r.Group("/")
	auth.Use(authMiddleware())

	auth.GET("/", message.IndexHandler)

	auth.GET("/api/messages", message.GetMessages)
	auth.POST("/api/messages", message.PostMessages)

	r.Run(":8888")
}
