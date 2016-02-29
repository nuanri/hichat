package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var _db *sql.DB

func GetConnection() *sql.DB {
	if _db == nil {
		_db = conn_database()
	}
	return _db
}

func conn_database() *sql.DB {
	db, err := sql.Open("mysql", "root:gui1gu2bai3nian4shi!@/hichatdb?charset=utf8")
	if err != nil {
		fmt.Println("Open database error: ", err)
	}

	//defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Ping error:", err)
	}
	return db
}
