package auth

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/hichatxyz/sendcloud"

	"nuanri/hichat/api/db"
)

func SignUpRequest(c *gin.Context) {
	s := &SignupStat{}
	err := c.BindJSON(s)
	if err == nil {
		conn := db.GetConnection()
		if s.verify_email_reg(conn) {
			log.Info("email already registered")
			return
		}
		s.Authcode = RandseqDigit(6)
		s.Authcode_key = Randseq(64)
		fmt.Println("注册邮箱是： ", s.Email)

		insert_authcode(conn, s.Authcode_key, s.Authcode, s.Email)

		SendMail(s.Email, s.Authcode)

		c.JSON(200, gin.H{"authcode_key": s.Authcode_key})
	} else {
		c.JSON(400, gin.H{"error": err.Error()})
	}

}

//使用 sendcloud 发送邮件
func SendMail(email string, authcode string) {
	s := sendcloud.SendCloud{
		ApiUser:  "hichat",
		ApiKey:   "20C0TGNizqxfnGQq",
		From:     "service@notice.hichat.xyz",
		Fromname: "HiChat 验证码",
	}

	tpl_name := "authcode_request"

	tos := []string{email}
	subs := map[string]interface{}{
		//邮件模板调用名称
		"%name%":     []string{email},
		"%authcode%": []string{authcode},
	}
	fmt.Println("name==>authocode", email, authcode)
	s.TemplateSend(tos, subs, tpl_name)
}

func SignUp(c *gin.Context) {
	s := &SignupStat{}

	if c.BindJSON(s) == nil {
		conn := db.GetConnection()
		if !s.verify_authcode(conn) {
			log.Error("verify_authcode return false")
			return
		}

		if s.verify_username_reg(conn) {
			log.Error("username already registered")
			return
		}

		if !insert_user(conn, s.Username, s.Password, s.Email) {
			log.Error("insert_user return false")
			return
		}
		c.JSON(200, gin.H{"authcode_key": s.Authcode_key, "authcode": s.Authcode, "username": s.Username, "password": s.Password})
	} else {
		fmt.Println("验证错误")
		c.JSON(400, gin.H{"error": "fail"})
		return
	}

}

func SignIn(c *gin.Context) {
	var l Login
	err := c.BindJSON(&l)
	if err == nil {
		if l.Username == "" {
			c.JSON(400, gin.H{"error": "no-username"})
			return
		}
		if l.Password == "" {
			c.JSON(400, gin.H{"error": "no-password"})
			return
		}
		conn := db.GetConnection()
		mark_username := l.verify_user(conn)
		if !mark_username {
			c.JSON(400, gin.H{"error": "密码或用户名错误！"})
			fmt.Println("Singin no find username")
			return
		} else {
			mark_password, user_id := l.user_login(conn)
			if !mark_password {
				c.JSON(400, gin.H{"error": "密码或用户名错误！"})
				fmt.Println("Singin password error")
				return
			} else {
				if !select_session_user_id(conn, user_id) {
					l.Sid = Randseq(128)
					insert_sid(conn, l.Sid, user_id)
				} else {
					l.Sid = select_sid(conn, user_id)
					fmt.Println("此用户的 sid 已存在!")
				}
			}
		}
		c.JSON(200, gin.H{"sid": l.Sid})

	} else {
		fmt.Println("form err:", err)
		c.JSON(400, gin.H{"error": err.Error()})
	}

}

func GetUserInfo(c *gin.Context) {
	conn := db.GetConnection()
	sid := c.Request.Header.Get("Sid")

	data := get_userinfo(conn, sid)
	c.JSON(200, gin.H{"id": data["id"], "email": data["email"], "username": data["username"], "last_activity_time": data["last_activity_time"]})

}

func Signout(c *gin.Context) {
	Sid, _ := c.Get("Sid")
	sid := Sid.(string)
	conn := db.GetConnection()
	signout_del_session(conn, sid)
}
