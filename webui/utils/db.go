package utils

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func OpenDB() *sql.DB {
	db, err := sql.Open("sqlite3", "hichat.db")
	//fmt.Println("sqlites==>", db)
	if err != nil {
		fmt.Println("Open sqlite3 database error: ", err)
	}
	return db
}

var _db *sqlx.DB

func InitDB() (err error) {
	if _db == nil {
		_db, err = sqlx.Connect("sqlite3", "hichat.db")
		if err != nil {
			return nil
		}
	}

	SQL, err := ioutil.ReadFile("utils/schema.sql")
	if err != nil {
		return err
	}

	fmt.Println("SQL = ", string(SQL))
	_db.MustExec(string(SQL))

	return nil
}
