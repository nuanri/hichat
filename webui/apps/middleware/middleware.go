package middleware

import (
	//log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"nuanri/hichat/webui/utils"
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
	db := utils.OpenDB()
	var u Userinfo
	cookie, err := c.Request.Cookie("Sid")
	if err != nil {
		return
	}

	session = &Session{
		Sid: cookie.Value,
	}

	u.get_userinfo(db, session.Sid)
	if !u.Online {
		u.set_useronline(db, u.Id)
	}

	session.User = u
	return
}
