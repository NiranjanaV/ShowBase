package tables

import (
	"fmt"
	"log"
	D "main/dbSQLite"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var tablename3 = D.GetTable(3)

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		log.Fatal("1 Error loading .env file" + err.Error())
	}
}

//************************************************************************************************************************************************************************
func CreateShowTable() {
	createUserTableSQL := `CREATE TABLE friend (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"idUser" integer NOT NULL,
		"idFriend" integer NOT NULL
	  );` // SQL Statement for Create User Table

	log.Println("Create user table friend")
	statement, err := db.Prepare(createUserTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("user table created")
}

//************************************************************************************************************************************************************************
// We are passing db reference connection from main to our method with other parameters
// func InsertUserTable(db *sql.DB, tablename string, user int, movie int, action int, value int) (error string) {
func InsertShowTable(c *gin.Context) {
	var msg string
	type UserFriendJson struct {
		Username  string `json:"username"`
		Username2 string `json:"friendname"`
	}
	userFriendJson := UserFriendJson{}
	err := c.ShouldBindJSON(&userFriendJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters, password should be between 8 to 20 chars",
		})
		return
	}

	user, _ := GetUserID(userFriendJson.Username)
	friend, err := GetUserID(userFriendJson.Username2)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	log.Println("Inserting student record ...")
	insertFriendSQL := `INSERT INTO ` + tablename3 + `(idUser, idFriend ) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertFriendSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec(user, friend)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Return": msg,
	})
}

//************************************************************************************************************************************************************************

type FriendNames struct {
	name string
}

var friendNames FriendNames

func GetShows(c *gin.Context) {

	idUser, _ := GetUserID(c.Param("username"))
	// fmt.Println(idUser)
	// fmt.Println("disp")
	row, err := db.Query("SELECT * FROM " + tablename3 + " WHERE idUser= " + strconv.Itoa(idUser))
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var id int
	var user int
	var friend int
	var friends []FriendNames
	for row.Next() { // Iterate and fetch the records from result cursor

		row.Scan(&id, &user, &friend)
		log.Println("Users: ", id, " ", user, " ", friend)
		friendNames.name = GetUserName(friend)
		friends = append(friends, friendNames)
	}
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"user":    user,
		"friends": friends,
	})
}

//************************************************************************************************************************************************************************
func DisplayShowTable(c *gin.Context) {
	fmt.Println("disp")
	row, err := db.Query("SELECT * FROM " + tablename2 + " ORDER BY idUser")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var id int
	var user int
	var friend int
	var friends []FriendNames
	for row.Next() { // Iterate and fetch the records from result cursor

		row.Scan(&id, &user, &friend)
		log.Println("Users: ", id, " ", user, " ", friend)
		friendNames.name = GetUserName(friend)
		friends = append(friends, friendNames)
	}
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"user":    user,
		"friends": friends,
	})
}

//************************************************************************************************************************************************************************
