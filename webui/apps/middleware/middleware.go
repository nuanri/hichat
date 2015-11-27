package middleware

import (
	"fmt"
	//log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"nuanri/hichat/webui/utils"
	"time"
)

type Userinfo struct {
	Id          int
	Email       string
	Username    string
	Online      bool
	LastMsgTime time.Time `json:"last_msg_time"`
	LastActTime time.Time `json:"last_act_time"`
}

type Session struct {
	Sid  string
	User Userinfo
}

func GetSession(c *gin.Context) (session *Session, err error) {
	db := utils.OpenDB()
	var u Userinfo
	cookie, err := c.Request.Cookie("Sid")
	fmt.Println("Sid-->", cookie)
	if err != nil {
		fmt.Println("hhhhhh--err", err)
		return nil, err
	}

	session = &Session{
		Sid: cookie.Value,
	}

	u.get_userinfo(db, session.Sid)
	//fmt.Println("----->", (time.Now().UTC().Format("2006-01-02 15:04:05")))
	if !u.Online {
		u.set_useronline(db, u.Id)
	}

	session.User = u
	return
}
