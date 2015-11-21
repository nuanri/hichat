package message

import (
	"fmt"
	//log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"html/template"

	"nuanri/hichat/webui/utils"
)

func IndexHandler(c *gin.Context) {

	var listTmpl = template.Must(template.ParseFiles("templates/base.html",
		"apps/message/templates/index.html"))

	db := utils.OpenDB()
	online_users := get_useronline(db)

	User, _ := c.Get("User")
	tc := make(map[string]interface{})
	tc["User"] = User
	tc["online_users"] = online_users

	if err := listTmpl.Execute(c.Writer, tc); err != nil {
		fmt.Println(err.Error())
	}
	//c.HTML(http.StatusOK, "auth_signup.html", gin.H{})
}

func GetMessages(c *gin.Context) {

}

func PostMessages(c *gin.Context) {

}
