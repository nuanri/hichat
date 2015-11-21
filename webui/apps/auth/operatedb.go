package auth

import (
	"database/sql"
	"fmt"
	//_ "github.com/mattn/go-sqlite3"
)

func insert_auth(db *sql.DB, user_id int, username string, password string, email string) {
	stmt, err := db.Prepare("insert into auth_user(id, username, password, email) values (?,?,?,?)")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	stmt.Exec(user_id, username, password, email)
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

/*
type Userinfo struct {
	Id       int
	Email    string
	Username string
}*/

/*func get_userinfo(db *sql.DB, sid string) interface{} {
	var u Userinfo
	row := db.QueryRow("select id, username, email from auth_user where id=(select user_id from auth_session where sid=?)", sid)
	err := row.Scan(&u.Id, &u.Username, &u.Email)
	if err != nil {
		fmt.Println("Signin Scan username failed:", err)
	}
	return u
}*/
