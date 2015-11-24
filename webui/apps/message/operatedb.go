package message

/*
import (
	"database/sql"
	"fmt"
)

func get_useronline(db *sql.DB) []string {
	var online_users []string

	rows, err := db.Query("select username from auth_user where online=1")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	var username string
	for rows.Next() {
		err = rows.Scan(&username)
		if err != nil {
			fmt.Println(err)
		}
		online_users = append(online_users, username)
	}
	return online_users
}
*/
