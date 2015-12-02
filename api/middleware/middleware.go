package middleware

import (
	//"fmt"
	//log "github.com/Sirupsen/logrus"
	"errors"
	"github.com/gin-gonic/gin"
	"nuanri/hichat/api/db"
	"time"
)

type Userinfo struct {
	Id          int
	Email       string
	Username    string
	Online      bool
	LastMsgTime string `json:"last_msg_time"`
	LastActTime string `json:"last_act_time"`
}

type Session struct {
	Sid  string
	User Userinfo
}

func GetSession(c *gin.Context) (session *Session, err error) {
	sid := c.Request.Header.Get("Sid")
	//fmt.Println("sid-->", sid)
	//fmt.Printf("Header = %#v\n", c.Request.Header)

	conn := db.GetConnection()
	var u Userinfo

	session = &Session{
		Sid: sid,
	}

	u.get_userinfo(conn, session.Sid)
	//fmt.Println("u===>", u)

	now := time.Now()
	utcnow := now.Add(-8 * time.Hour)
	d_time := utcnow.Add(-60 * time.Minute)
	l_time := d_time.Format("2006-01-02 15:04:05")

	//fmt.Println("u.LastActTime == >", u.LastActTime)
	if (u.LastActTime != "") && (l_time > u.LastActTime) {
		u.delete_sid(conn, u.Id)
		return nil, errors.New("session expired")
	}

	if !u.Online {
		u.set_useronline(conn, u.Id)
	}

	session.User = u

	return
}
