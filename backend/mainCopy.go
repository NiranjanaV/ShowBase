package main

import (
	DM "main/functions/dbAccess"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//************************************************************************************************************************************************************************

func main() {

	r.PUT("/userReg", DM.InsertAuthTable)
}

//************************************************************************************************************************************************************************
