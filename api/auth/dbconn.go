package auth

import (
	"database/sql"
	"fmt"
	//log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

type SignupStat struct {
	Authcode     string `json:"authcode"`
	Authcode_key string `json:"authcode_key"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

//注册填写email时先检测是否邮箱已注册
func (s *SignupStat) verify_email_reg(db *sql.DB) bool {
	row := db.QueryRow("select count(id) from auth_user  where email=?", s.Email)
	var count int
	err := row.Scan(&count)
	if err != nil {
		fmt.Println("Signuprequest reg scan  email failed:", err)
	}
	if count > 0 {
		return true
	}
	return false
}

//判断 user_name 是否被用
func (s *SignupStat) verify_username_reg(db *sql.DB) bool {
	row := db.QueryRow("select count(id) from auth_user  where username=?", s.Username)
	var count int
	err := row.Scan(&count)
	if err != nil {
		fmt.Println("Signuprequest reg scan  username failed:", err)
	}
	if count > 0 {
		return true
	}
	return false
}

func (s *SignupStat) verify_authcode(db *sql.DB) bool {
	rows, err := db.Query("select id from auth_code where authcode=? and authcode_key=?", s.Authcode, s.Authcode_key)
	if err != nil {
		fmt.Println("Query failed:", err)
		return false
	}
	defer rows.Close()
	var id int
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println("Scan failed:", err)
			return false
		}
		return true
	}
	return false
}

func insert_authcode(db *sql.DB, authcode_key string, authcode string, email string) {
	row := db.QueryRow("select count(id) from auth_code  where email=?", email)
	var count int
	err := row.Scan(&count)
	if err != nil {
		fmt.Println("Signuprequest Scan email failed:", err)
		return
	}
	if count > 0 {
		stmt, err := db.Prepare("update auth_code set authcode_key=?, authcode=? where email=?")
		if err != nil {
			fmt.Println(err)
			return
		}
		stmt.Exec(authcode_key, authcode, email)
		defer stmt.Close()
		return
	}

	stmt, err := db.Prepare("insert into auth_code(authcode_key, authcode, email) VALUES(?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	stmt.Exec(authcode_key, authcode, email)
}

func insert_user(db *sql.DB, username string, password string, email string) bool {
	stmt, err := db.Prepare("insert into auth_user(username, password, email) VALUES(?, ?, ?)")
	if err != nil {
		return false
	}
	defer stmt.Close()
	stmt.Exec(username, password, email)
	return true
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Sid      string `json:"sid"`
}

//验证此用户是否存在 auth_user
func (l *Login) verify_user(db *sql.DB) bool {
	mark_username := false

	row := db.QueryRow("select id from auth_user  where username=?", l.Username)
	var id int
	err := row.Scan(&id)
	if err != nil {
		fmt.Println("Signin Scan username failed:", err)
		return mark_username
	}
	mark_username = true
	return mark_username
}

//验证 password
func (l *Login) user_login(db *sql.DB) (bool, int) {
	mark_password := false

	row := db.QueryRow("select id from auth_user where password=? and username=?", l.Password, l.Username)
	var id int
	err := row.Scan(&id)
	if err != nil {
		fmt.Println("login Scan password failed:", err)
		return mark_password, id
	}
	mark_password = true
	return mark_password, id
}

//查看  auth_session 中 user_id 是否存在
func select_session_user_id(db *sql.DB, user_id int) bool {
	row := db.QueryRow("select id from auth_session where user_id=?", user_id)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return false
	} else {
		return true
	}

}

//向 auth_session 中插入记录
func insert_sid(db *sql.DB, sid string, user_id int) {

	stmt, err := db.Prepare("insert into auth_session(user_id, sid) VALUES(?, ?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	stmt.Exec(user_id, sid)
}

//如果  user_id 在 auth_session 存在，就去当前的 sid
func select_sid(db *sql.DB, user_id int) string {

	row := db.QueryRow("select sid from auth_session where user_id=?", user_id)
	var sid string
	err := row.Scan(&sid)
	if err != nil {
		fmt.Println(err)
		return sid
	} else {
		return sid
	}

}

type UserInfo struct {
	Sid string `json:"sid"`
}

func get_userinfo(db *sql.DB, sid string) map[string]interface{} {

	row := db.QueryRow("select a.id, a.username, a.email, a.last_activity_time from  auth_user a, auth_session b where a.id=b.user_id and b.sid=?", sid)
	var id int
	var username string
	var last_activity_time string
	var email string

	err := row.Scan(&id, &username, &email, &last_activity_time)
	if err != nil {
		fmt.Println(err)
	}

	data := map[string]interface{}{
		"id":                 id,
		"username":           username,
		"email":              email,
		"last_activity_time": last_activity_time,
	}
	//fmt.Println("data===>", data)
	return data
}

func signout_del_session(db *sql.DB, sid string) {
	stmt, err := db.Prepare(`DELETE FROM auth_session WHERE sid=?`)
	if err != nil {
		fmt.Println(err)
		return
	}

	stmt.Exec(sid)
}
