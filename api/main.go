package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"nuanri/hichat/api/auth"
	"nuanri/hichat/api/message"
	"nuanri/hichat/api/middleware"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		session, err := middleware.GetSession(c)
		if err != nil {
			fmt.Println("===>", err)
			c.String(400, err.Error())
			c.Abort()
			return
		}

		if session != nil {
			// Set example variable
			c.Set("Sid", session.Sid)
			c.Set("User", session.User)
		}
		// before request
		c.Next()
		// after request
	}
}

func main() {

	router := gin.Default()
	router.Use(CORSMiddleware())
	router.Use(SessionMiddleware())

	// auth
	router.POST("/signup/request", auth.SignUpRequest)
	router.POST("/register/passwd", auth.SignUp)
	router.POST("/auth/signin", auth.SignIn)
	router.GET("/auth/userinfo", auth.GetUserInfo)

	// message
	router.GET("/messages", message.GetMessages)
	router.POST("/messages", message.NewMessage)

	router.Run(":8080")
}
