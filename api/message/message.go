package message

import (
	"fmt"
	"net/http"
	"time"
	//log "github.com/Sirupsen/logrus"
	//"database/sql"
	"github.com/gin-gonic/gin"
	//_ "github.com/go-sql-driver/mysql"

	"nuanri/hichat/api/db"
	"nuanri/hichat/api/middleware"
)

// Binding from JSON
type Message struct {
	Body string `json:"body" binding:"required"`
	To   string `json:"to"`
}

func NewMessage(c *gin.Context) {
	user, _ := c.Get("User")

	username := user.(middleware.Userinfo).Username
	//sid, _ := c.Get("Sid")

	var msg Message
	if c.BindJSON(&msg) == nil {
		//fmt.Printf("%#v\n", msg)
		conn := db.GetConnection()
		insert_message(conn, msg.Body, username)
		if msg.To != "" {
			//fmt.Println("go to:", msg.To)
			c.JSON(http.StatusOK, gin.H{"status": "have to", "body": "you say: " + msg.Body})
		} else {
			//fmt.Println("to all")
			c.JSON(http.StatusOK, gin.H{"status": "to all", "body": "you say: " + msg.Body})
		}
		return
	}

	c.JSON(400, gin.H{"error": "system-error"})
}

// 查询新消息
func GetMessages(c *gin.Context) {
	//user, _ := c.Get("User")
	//sid, _ := c.Get("Sid")
	//fmt.Println("get user===>sid", user, sid)
	var data_time []map[string]interface{}
	var online_users []string

	conn := db.GetConnection()

	if lasttime := c.Query("t"); len(lasttime) > 0 {
		lt_dt, _ := time.Parse(time.RFC3339Nano, lasttime)
		mysql_dt := lt_dt.Format("2006-01-02 15:04:05")

		now := time.Now()
		utcnow := now.Add(-8 * time.Hour)
		d_time := utcnow.Add(-10 * time.Minute)
		drop_time := d_time.Format("2006-01-02 15:04:05")
		fmt.Println("drop_time", drop_time)

		for {
			data_time = select_message_time(conn, mysql_dt)
			//fmt.Println(data_time)
			if len(data_time) > 0 {
				break
			}
			time.Sleep(1 * time.Second)
		}

		online_users = get_useronline(conn, drop_time)

	} else {

		now := time.Now()
		utcnow := now.Add(-8 * time.Hour)
		d_time := utcnow.Add(-10 * time.Minute)
		drop_time := d_time.Format("2006-01-02 15:04:05")

		data_time = select_message_new(conn)
		online_users = get_useronline(conn, drop_time)

	}

	c.JSON(http.StatusOK, gin.H{"body": data_time, "onlineusers": online_users})
}
