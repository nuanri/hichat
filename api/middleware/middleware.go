package middleware

import (
	//"fmt"
	//log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"nuanri/hichat/api/db"
)

type Userinfo struct {
	Id       int
	Email    string
	Username string
	Online   bool
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
	if !u.Online {
		u.set_useronline(conn, u.Id)
	}

	session.User = u

	return
}
