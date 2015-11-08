package message

import (
	"fmt"
	"net/http"
	"time"
	//log "github.com/Sirupsen/logrus"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"nuanri/hichat/api/db"
)

// Binding from JSON
type Message struct {
	Body string `json:"body" binding:"required"`
	To   string `json:"to"`
}

func NewMessage(c *gin.Context) {
	fmt.Println("got: ", c)
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

func insert_message(db *sql.DB, body string) {
	stmt, err := db.Prepare("insert into msg_record(msg, add_time) VALUES(?,?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	stmt.Exec(body, time.Now())
}

func select_message_first(db *sql.DB) []map[string]interface{} {
	rows, err := db.Query("select msg, add_time from msg_record limit 10;")
	if err != nil {
		fmt.Println("err1:", err)
	}

	data := new([]map[string]interface{})
	for rows.Next() {
		var msg string
		var add_time string
		rows.Columns()
		err = rows.Scan(&msg, &add_time)
		if err != nil {
			fmt.Println("err2:", err)
		}
		*data = append(*data, map[string]interface{}{
			"msg":      msg,
			"add_time": add_time,
		})

	}
	return *data
}

func select_message_time(db *sql.DB, mysql_dt string) []map[string]interface{} {
	rows, err := db.Query("select msg, add_time from msg_record where add_time >=?", mysql_dt)
	if err != nil {
		fmt.Println("err3:", err)
	}
	defer rows.Close()
	data := new([]map[string]interface{})
	for rows.Next() {
		var msg string
		var add_time string
		rows.Columns()
		err = rows.Scan(&msg, &add_time)
		fmt.Println("find: ", add_time, msg)
		if err != nil {
			fmt.Println("err4:", err)
		}
		*data = append(*data, map[string]interface{}{
			"msg":      msg,
			"add_time": add_time,
		})

	}
	//fmt.Println(*data)
	return *data
}

// 查询新消息
func GetMessages(c *gin.Context) {
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
