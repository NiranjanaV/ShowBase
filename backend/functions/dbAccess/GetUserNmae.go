package tables

import (
	"fmt"
	"log"
	D "main/dbSQLite"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		log.Fatal("1 Error loading .env file" + err.Error())
	}
}

//************************************************************************************************************************************************************************
func GetUserName(useridInt int) (error string) {
	// func InsertUserTable(c *gin.Context) {
	log.Println("getting user ", useridInt)
	row, err := db.Query("SELECT username FROM " + D.GetTable(1) + " WHERE userID= '" + strconv.Itoa(useridInt) + "'")

	if err != nil {
		fmt.Println(err.Error())
	}
	//defer
	for row.Next() { // Iterate and fetch the records from result cursor
		var usernameRet string
		row.Scan(&usernameRet)
		row.Close()
		fmt.Println(usernameRet)
		error = usernameRet
	}
	return
}

//************************************************************************************************************************************************************************
