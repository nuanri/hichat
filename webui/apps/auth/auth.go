package auth

import (
	"fmt"
	//	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"

	"nuanri/hichat/webui/utils"
)

func SignupRequest(c *gin.Context) {

	var listTmpl = template.Must(template.ParseFiles("templates/base.html",
		"apps/auth/templates/auth_signup.html"))
	tc := make(map[string]interface{})
	//tc["Store"] = "ccccc"
	//tc["Products"] = "bbbb"

	if err := listTmpl.Execute(c.Writer, tc); err != nil {
		fmt.Println(err.Error())
	}
}

func GetSignin(c *gin.Context) {

	var listTmpl = template.Must(template.ParseFiles("templates/base.html",
		"apps/auth/templates/auth_signin.html"))
	tc := make(map[string]interface{})

	if err := listTmpl.Execute(c.Writer, tc); err != nil {
		fmt.Println(err.Error())
	}

	//c.HTML(http.StatusOK, "auth_signin.html", gin.H{})
}

type LS struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type M struct {
	Sid string `json:"sid"`
}

type UserInfo struct {
	Userid   int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func PostSignin(c *gin.Context) {
	var form LS
	err := c.Bind(&form)
	if err == nil {
		b := GetJson(form)
		url := "http://192.168.0.7:8080/auth/signin"
		method := "POST"
		body := GetBackendApi("", method, url, b)
		fmt.Println("body=", body)
		var m M
		ParseMsg(body, &m)
		// 设置 web client cookie
		http.SetCookie(c.Writer, &http.Cookie{Name: "Sid", Value: m.Sid, Path: "/"})

		user_url := "http://192.168.0.7:8080/auth/userinfo"
		user_method := "GET"
		uinfo := GetBackendApi(m.Sid, user_method, user_url, nil)

		var u UserInfo
		ParseMsg(uinfo, &u)
		fmt.Printf("u = %#v\n", u)
		conn := utils.OpenDB()
		insert_auth(conn, u.Userid, u.Username, u.Password, u.Email)
		insert_session(conn, u.Userid, m.Sid)

		c.Redirect(302, "/")
	}
	fmt.Println("bind failed:", err)

}
