package auth

import (
	"database/sql"
	"fmt"
	//_ "github.com/mattn/go-sqlite3"
)

func insert_auth(db *sql.DB, user_id int, username string, last_activity_time string, email string) {

	row := db.QueryRow("select count(id) from auth_user  where user_id=?", user_id)
	var count int
	err := row.Scan(&count)
	if err != nil {
		fmt.Println(err)
	}
	if count > 0 {
		return
	}

	stmt, err := db.Prepare("insert into auth_user(user_id, username, email, last_activity_time) values (?,?,?,?)")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	stmt.Exec(user_id, username, email, last_activity_time)
}

func insert_session(db *sql.DB, user_id int, sid string) {
	row := db.QueryRow("select count(id) from auth_session  where user_id=?", user_id)
	var count int
	err := row.Scan(&count)
	if err != nil {
		fmt.Println(err)
	}
	if count > 0 {
		return
	}
	stmt, err := db.Prepare("insert into auth_session(user_id, sid) values (?,?)")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	if _, err := stmt.Exec(user_id, sid); err != nil {
		fmt.Println("smt.Exec failed: ", err)
	}
}

func signout_del_session(db *sql.DB, sid string) {
	stmt, err := db.Prepare(`DELETE FROM auth_session WHERE sid=?`)
	if err != nil {
		fmt.Println(err)
		return
	}

	stmt.Exec(sid)
}
