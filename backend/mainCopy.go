package main

import (
	// D "main/dbSQLite"
	// A "main/functions/APIcalls"
	// DM "main/functions/dbAccess"
	DM "main/functions/dbAccess"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//************************************************************************************************************************************************************************

func main() {
	DM.CreateFriendTable()
}

//************************************************************************************************************************************************************************
