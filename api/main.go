package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"nuanri/hichat/api/auth"
	"nuanri/hichat/api/message"
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

func main() {

	router := gin.Default()
	router.Use(CORSMiddleware())

	// auth
	router.POST("/signup/request", auth.SignUpRequest)
	router.POST("/register/passwd", auth.SignUp)
	router.POST("/auth/signin", auth.SignIn)

	// message
	router.GET("/messages", message.GetMessages)
	router.POST("/messages", message.NewMessage)

	router.Run(":8080")
}
