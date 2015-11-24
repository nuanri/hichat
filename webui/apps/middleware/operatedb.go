package middleware

import (
	"database/sql"
	"fmt"
	"time"
	//_ "github.com/mattn/go-sqlite3"
)

func (u *Userinfo) get_userinfo(db *sql.DB, sid string) {

	row := db.QueryRow("select user_id, username, email, online from auth_user where user_id=(select user_id from auth_session where sid=?)", sid)
	err := row.Scan(&u.Id, &u.Username, &u.Email, &u.Online)
	if err != nil {
		fmt.Println("Signin Scan username failed:", err)
	}
}

func (u *Userinfo) set_useronline(db *sql.DB, id int) {

	stmt, err := db.Prepare(`UPDATE auth_user SET online=?, last_activity_time=?  WHERE user_id=?`)
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec(1, time.Now(), id)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
	}
}
