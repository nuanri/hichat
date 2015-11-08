package auth

import (
	"fmt"

	log "github.com/Sirupsen/logrus"

	"github.com/gin-gonic/gin"

	"nuanri/hichat/api/db"
)

func SignUpRequest(c *gin.Context) {
	s := &SignupStat{}

	if c.BindJSON(s) == nil {
		s.Authcode = RandseqDigit(6)
		s.Authcode_key = Randseq(64)

		fmt.Println("注册邮箱是： ", s.Email)
		fmt.Println(s.Authcode, "==>", s.Authcode_key)

		conn := db.GetConnection()
		insert_authcode(conn, s.Authcode_key, s.Authcode, s.Email)

		c.JSON(200, gin.H{"authcode_key": s.Authcode_key})
	} else {
		c.JSON(400, gin.H{"error": "fail"})
	}

}

func SignUp(c *gin.Context) {
	s := &SignupStat{}

	if c.BindJSON(s) == nil {
		fmt.Println("s = ", s)
		conn := db.GetConnection()
		if !s.verify_authcode(conn) {
			fmt.Println("===========>")
			log.Error("verify_authcode return false")
			return
		}
		if !insert_user(conn, s.Username, s.Password) {
			log.Error("insert_user return false")
			return
		}
		c.JSON(200, gin.H{"authcode_key": s.Authcode_key, "authcode": s.Authcode, "username": s.Username, "password": s.Password})
	} else {
		fmt.Println("验证错误")
		c.JSON(400, gin.H{"erroe": "fail"})
		return
	}

}
