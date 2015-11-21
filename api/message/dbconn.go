package message

import (
	"database/sql"
	"fmt"
	"time"
	//log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

func insert_message(db *sql.DB, body string) {
	stmt, err := db.Prepare("insert into msg_record(msg, add_time) VALUES(?,?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	stmt.Exec(body, time.Now())
}

/*
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
*/

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

/*
func get_useronline(db *sql.DB) []string {
	var online_users []string

	rows, err := db.Query("select username from auth_user where online=1")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	var username string
	for rows.Next() {
		err = rows.Scan(&username)
		if err != nil {
			fmt.Println(err)
		}
		online_users = append(online_users, username)
	}
	return online_users
}
*/
