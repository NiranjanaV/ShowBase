package dataBasePkg

import (
	_ "github.com/mattn/go-sqlite3"
)

func GetTable(tablenumber int) (tablename string) {

	//returning table name
	if tablenumber == 1 {
		tablename = "user_auth"
	} else if tablenumber == 2 {
		tablename = "user"
	}
	return
}
