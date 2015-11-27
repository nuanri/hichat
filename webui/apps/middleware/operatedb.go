package middleware

import (
	"database/sql"
	"fmt"
	//_ "github.com/mattn/go-sqlite3"
)

func (u *Userinfo) get_userinfo(db *sql.DB, sid string) {
	fmt.Println("--->--->", sid)

	stmt, err := db.Prepare("UPDATE auth_user SET last_act_time=datetime('now')  WHERE id=(select user_id from auth_session where sid=?)")
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec(sid)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
	}

	row := db.QueryRow("select user_id, username, email, online, last_msg_time, last_act_time from auth_user where user_id=(select user_id from auth_session where sid=?)", sid)
	err = row.Scan(&u.Id, &u.Username, &u.Email, &u.Online, &u.LastMsgTime, &u.LastActTime)
	if err != nil {
		fmt.Println("===>Signin Scan username failed:", err)
	}
}

func (u *Userinfo) set_useronline(db *sql.DB, id int) {
	//lmt := time.Now().UTC().Format("2006-01-02 15:04:05")
	//fmt.Println("++++>lmt", lmt)

	stmt, err := db.Prepare("UPDATE auth_user SET online=?, last_msg_time=datetime('now') WHERE user_id=?")
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec(1, id)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
	}
}
