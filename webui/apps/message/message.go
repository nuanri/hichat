package message

import (
	"fmt"
	//log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"html/template"
	//"net/http"

	"nuanri/hichat/webui/apps/auth"
)

func IndexHandler(c *gin.Context) {

	var listTmpl = template.Must(template.ParseFiles("templates/base.html",
		"apps/message/templates/index.html"))

	User, _ := c.Get("User")
	fmt.Println("User+++==>", User)

	bapi := auth.GetBackendApi2(c)
	data := gin.H{}
	if err := bapi.Get(&data, "http://192.168.0.7:8080/messages"); err != nil {
		fmt.Println("bapi failed:", err)
		return
	}
	data["User"] = User

	if err := listTmpl.Execute(c.Writer, data); err != nil {
		fmt.Println(err.Error())
	}

}

func GetMessages(c *gin.Context) {

	cookie, err := c.Request.Cookie("Sid")
	if err != nil {
		fmt.Println(err)
	}
	sid := cookie.Value
	lasttime := c.Query("t")
	url := "http://192.168.0.7:8080/messages?t=" + lasttime
	method := "GET"
	var b []byte
	body := GetBackendApi(sid, method, url, b)
	//fmt.Printf("get messages: body = %#v\n", body)
	c.String(200, string(body))
}

type SMessage struct {
	To   string
	Body string
}

func PostMessages(c *gin.Context) {
	cookie, err := c.Request.Cookie("Sid")
	if err != nil {
		fmt.Println(err)
	}
	sid := cookie.Value

	var msg SMessage
	err = c.Bind(&msg)
	if err == nil {
		b := GetJson(msg)
		url := "http://192.168.0.7:8080/messages"
		method := "POST"
		body := GetBackendApi(sid, method, url, b)
		//fmt.Printf("body = %#v\n", string(body))
		c.String(200, string(body))
	}
	return
}
