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
	db, err := sql.Open("mysql", "root:abc@/hichartdb")
	fmt.Println("db--->", db)
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
