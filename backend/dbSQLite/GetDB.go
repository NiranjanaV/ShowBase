package dataBasePkg

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB() (db *sql.DB) {

	// Open the created SQLite File
	db, _ = sql.Open("sqlite3", "dbSQLite/user_data-database")

	return
}
