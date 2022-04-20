package dataBasePkg

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB() (db *sql.DB) {
	var err error
	// Open the created SQLite File
	db, err = sql.Open("sqlite3", "dbSQLite/user_data-database")
	// //fmt.Println(db.Query("Select count(*) from user"))
	if err != nil {
		//fmt.Println("Issue here itsrelf ")
	}
	return
}
