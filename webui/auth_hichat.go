package main

import (
	//"database/sql"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	//_ "github.com/go-sql-driver/mysql"

	//"net/http"
)

func main() {

	db := conn_database()
	router := gin.Default()

	router.POST("/signup/request", func(c *gin.Context) {
		s := &SignupStat{}

		if c.BindJSON(s) == nil {
			s.Authcode = RandseqDigit(6)
			s.Authcode_key = Randseq(64)

			fmt.Println("注册邮箱是： ", s.Email)
			fmt.Println(s.Authcode, "==>", s.Authcode_key)

			insert_authcode(db, s.Authcode_key, s.Authcode, s.Email)

			c.JSON(200, gin.H{"authcode_key": s.Authcode_key})
		} else {
			c.JSON(400, gin.H{"error": "fail"})
		}

	})

	router.POST("/register/passwd", func(c *gin.Context) {
		s := &SignupStat{}

		if c.BindJSON(s) == nil {
			fmt.Println("s = ", s)
			if !s.verify_authcode(db) {
				fmt.Println("===========>")
				log.Error("verify_authcode return false")
				return
			}
			if !insert_user(db, s.Username, s.Password) {
				log.Error("insert_user return false")
				return
			}
			fmt.Println("222")
			c.JSON(200, gin.H{"authcode_key": s.Authcode_key, "authcode": s.Authcode, "username": s.Username, "password": s.Password})
		} else {
			fmt.Println("验证错误")
			c.JSON(400, gin.H{"erroe": "fail"})
			return
		}

	})

	router.Run(":8080")
}
