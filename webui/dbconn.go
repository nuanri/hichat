package main

import (
	"database/sql"
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

type SignupStat struct {
	Authcode     string `json:"authcode"`
	Authcode_key string `json:"authcode_key"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

func (s *SignupStat) verify_authcode(db *sql.DB) bool {
	rows, err := db.Query("select id from authcode where authcode=? and authcode_key=?", s.Authcode, s.Authcode_key)
	if err != nil {
		log.Error("Query failed:", err)
		return false
	}
	defer rows.Close()
	var id int
	for rows.Next() {
		fmt.Println("b id=>", id)
		err = rows.Scan(&id)
		fmt.Println("a id=>", id)
		if err != nil {
			log.Error("Scan failed:", err)
			return false
		}
		return true
	}
	return false
}

func insert_authcode(db *sql.DB, authcode_key string, authcode string, email string) {
	stmt, err := db.Prepare("insert into authcode(authcode_key, authcode, email) VALUES(?, ?, ?)")
	if err != nil {
		log.Error(err)
		return
	}
	defer stmt.Close()
	stmt.Exec(authcode_key, authcode, email)
}

func insert_user(db *sql.DB, username string, password string) bool {
	stmt, err := db.Prepare("insert into auth_user(username, password) VALUES(?, ?)")
	if err != nil {
		log.Error("inset_user error=>", err)
		return false
	}
	defer stmt.Close()
	stmt.Exec(username, password)
	return true
}

func conn_database() *sql.DB {
	db, err := sql.Open("mysql", "root:abc@/hichartdb")
	fmt.Println("db--->", db)
	if err != nil {
		log.Error("Open database error: ", err)
	}

	//defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Error("Ping error:", err)
	}
	return db
}
