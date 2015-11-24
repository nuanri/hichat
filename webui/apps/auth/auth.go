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
	Sid   string `json:"sid"`
	Error string
}

type UserInfo struct {
	Userid             int `json:"id"`
	Username           string
	Last_activity_time string
	Email              string
}

func PostSignin(c *gin.Context) {
	var listTmpl = template.Must(template.ParseFiles("templates/base.html",
		"apps/auth/templates/auth_signin.html"))

	var form LS
	err := c.Bind(&form)
	if err == nil {
		b := GetJson(form)
		url := "http://192.168.0.7:8080/auth/signin"
		method := "POST"
		body := GetBackendApi("", method, url, b)

		var m M
		ParseMsg(body, &m)
		signinerror := m.Error
		if signinerror != "" {
			tc := make(map[string]interface{})
			tc["SError"] = signinerror

			if err := listTmpl.Execute(c.Writer, tc); err != nil {
				fmt.Println(err.Error())
			}
			return
		}
		// 设置 web client cookie
		http.SetCookie(c.Writer, &http.Cookie{Name: "Sid", Value: m.Sid, Path: "/"})
		bapi := GetBackendApi2(c)
		bapi.Sid = m.Sid
		u := &UserInfo{}
		if err := bapi.Get(&u, "http://192.168.0.7:8080/auth/userinfo"); err != nil {
			fmt.Println("获取用户信息出错:", err)
			return
		}

		//fmt.Printf("u = %#v\n", u)
		conn := utils.OpenDB()
		insert_auth(conn, u.Userid, u.Username, u.Last_activity_time, u.Email)
		insert_session(conn, u.Userid, m.Sid)

		c.Redirect(302, "/")
	}
	fmt.Println("bind failed:", err)

}

func Siginout(c *gin.Context) {
	conn := utils.OpenDB()

	cookie, err := c.Request.Cookie("Sid")
	if err != nil {
		fmt.Println("Signout 出错！")
		return
	}
	sid := cookie.Value
	signout_del_session(conn, sid)
	url := "http://192.168.0.7:8080/auth/signout"
	method := "GET"
	var b []byte
	body := GetBackendApi(sid, method, url, b)
	fmt.Println("1111body=", string(body))
	sid = ""
	// 设置 web client cookie
	http.SetCookie(c.Writer, &http.Cookie{Name: "Sid", Value: sid, Path: "/"})
}
