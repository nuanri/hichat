package middleware

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func (u *Userinfo) get_userinfo(db *sql.DB, sid string) {

	row := db.QueryRow("select id, username, email, online, last_msg_time, last_act_time from auth_user where id=(select user_id from auth_session where sid=?)", sid)
	err := row.Scan(&u.Id, &u.Username, &u.Email, &u.Online, &u.LastMsgTime, &u.LastActTime)
	if err != nil {
		fmt.Println("Signin Scan username failed:", err)
	}

	stmt, err := db.Prepare(`UPDATE auth_user SET last_act_time=?  WHERE id=(select user_id from auth_session where sid=?)`)
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec(time.Now(), sid)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
	}
}

func (u *Userinfo) set_useronline(db *sql.DB, id int) {

	stmt, err := db.Prepare("UPDATE auth_user SET online=?, last_msg_time=?  WHERE id=?")
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec(1, time.Now(), id)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
	}
}

func (u *Userinfo) delete_sid(db *sql.DB, id int) {

	stmt, err := db.Prepare(`DELETE FROM auth_session WHERE user_id=?`)
	if err != nil {
		fmt.Println(err)
		return
	}

	stmt.Exec(id)

}
