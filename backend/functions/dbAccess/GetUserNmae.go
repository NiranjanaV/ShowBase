package tables

import (
	D "main/dbSQLite"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		//fmt.Println("1 Error loading .env file" + err.Error())
	}
}

//************************************************************************************************************************************************************************
func GetUserName(useridInt int) (name string) {
	// func InsertUserTable(c *gin.Context) {
	//log.Println("getting user ", useridInt)
	row, err := db.Query("SELECT username FROM " + D.GetTable(1) + " WHERE idUser= '" + strconv.Itoa(useridInt) + "'")

	if err != nil {
		//fmt.Println(err.Error())
	}
	//defer
	for row.Next() { // Iterate and fetch the records from result cursor
		var usernameRet string
		row.Scan(&usernameRet)
		row.Close()
		//fmt.Println(usernameRet)
		name = usernameRet
	}
	return
}

//************************************************************************************************************************************************************************
