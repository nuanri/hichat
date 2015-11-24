package message

import (
	"database/sql"
	"fmt"
	"time"
	//log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

//存入消息，并 update auth_user 表里的 last_activity_time
func insert_message(db *sql.DB, body string, username string) {

	stmt, err := db.Prepare("insert into msg_record(msg, user, add_time) VALUES(?,?,?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	stmt.Exec(body, username, time.Now())

	stmt, err = db.Prepare("update auth_user set last_activity_time=? where username=?")
	if err != nil {
		fmt.Println(err)
		return
	}
	stmt.Exec(time.Now(), username)
	defer stmt.Close()
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
	rows, err := db.Query("select msg, user, add_time from msg_record where add_time >=?", mysql_dt)
	if err != nil {
		fmt.Println("err3:", err)
	}
	defer rows.Close()
	data := new([]map[string]interface{})
	for rows.Next() {
		var msg string
		var user string
		var add_time string
		rows.Columns()
		err = rows.Scan(&msg, &user, &add_time)
		if err != nil {
			fmt.Println("err4:", err)
		}
		*data = append(*data, map[string]interface{}{
			"msg":      msg,
			"add_time": add_time,
			"username": user,
		})

	}

	return *data
}

func select_message_new(db *sql.DB) []map[string]interface{} {
	rows, err := db.Query("select msg, user, add_time from  (select * from msg_record order by id desc limit 20) as T1 order by id")
	if err != nil {
		fmt.Println("err3:", err)
	}
	defer rows.Close()
	data := new([]map[string]interface{})
	for rows.Next() {
		var msg string
		var user string
		var add_time string
		rows.Columns()
		err = rows.Scan(&msg, &user, &add_time)
		if err != nil {
			fmt.Println("err4:", err)
		}
		*data = append(*data, map[string]interface{}{
			"msg":      msg,
			"add_time": add_time,
			"username": user,
		})

	}

	return *data
}

func get_useronline(db *sql.DB, drop_time string) []string {
	var online_users []string

	stmt, err := db.Prepare(`UPDATE auth_user SET online=?  WHERE last_activity_time<=?`)
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec(0, drop_time)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
	}

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
