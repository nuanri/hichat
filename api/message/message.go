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
	//"nuanri/hichat/webui/utils"
)

// Binding from JSON
type Message struct {
	Body string `json:"body" binding:"required"`
	To   string `json:"to"`
}

func NewMessage(c *gin.Context) {
	user, _ := c.Get("User")
	sid, _ := c.Get("Sid")
	fmt.Println("send user===>sid", user, sid)
	var msg Message
	if c.BindJSON(&msg) == nil {
		fmt.Printf("%#v\n", msg)
		conn := db.GetConnection()
		insert_message(conn, msg.Body)
		if msg.To != "" {
			fmt.Println("go to:", msg.To)
			c.JSON(http.StatusOK, gin.H{"status": "have to", "body": "you say: " + msg.Body})
		} else {
			fmt.Println("to all")
			c.JSON(http.StatusOK, gin.H{"status": "to all", "body": "you say: " + msg.Body})
		}
		return
	}

	c.JSON(400, gin.H{"error": "system-error"})
}

// 查询新消息
func GetMessages(c *gin.Context) {
	user, _ := c.Get("User")
	sid, _ := c.Get("Sid")
	fmt.Println("get user===>sid", user, sid)

	conn := db.GetConnection()
	//data := select_message_first(conn)
	lasttime := c.Query("t")
	//lt_int, _ := strconv.ParseInt(lasttime, 10, 64)
	lt_dt, _ := time.Parse(time.RFC3339Nano, lasttime)
	mysql_dt := lt_dt.Format("2006-01-02 15:04:05")
	var data_time []map[string]interface{}

	for {
		data_time = select_message_time(conn, mysql_dt)
		//fmt.Println(data_time)
		if len(data_time) > 0 {
			break
		}
		time.Sleep(1 * time.Second)
	}
	c.JSON(http.StatusOK, gin.H{"body": data_time})
}
