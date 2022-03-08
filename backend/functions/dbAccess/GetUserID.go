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
func GetUserID(usernameStr string) (userIdGet int, errR error) {
	// func InsertUserTable(c *gin.Context) {
	log.Println("getting user ", usernameStr)
	row, err := db.Query("SELECT idUser FROM " + D.GetTable(1) + " WHERE username= '" + usernameStr + "'")

	if err != nil {
		fmt.Println(err.Error())
		errR = err
	}
	//defer
	for row.Next() { // Iterate and fetch the records from result cursor
		var userId string
		row.Scan(&userId)
		row.Close()
		fmt.Println(userId)
		userIdGet, errR = strconv.Atoi(userId)
		fmt.Println(errR.Error())
	}
	return
}

//************************************************************************************************************************************************************************
