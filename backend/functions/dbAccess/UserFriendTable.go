package tables

import (
	D "main/dbSQLite"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var tablename3 = D.GetTable(3)
var tablename1 = D.GetTable(1)

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		//fmt.Println("1 Error loading .env file" + err.Error())
	}
}

//************************************************************************************************************************************************************************
func CreateFriendTable() {
	createUserTableSQL := `CREATE TABLE friend (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"idUser" integer NOT NULL,
		"idFriend" integer NOT NULL
	  );` // SQL Statement for Create User Table

	//log.Println("Create user table friend")
	statement, err := db.Prepare(createUserTableSQL) // Prepare SQL Statement
	if err != nil {
		//fmt.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	//log.Println("user table created")
}

//************************************************************************************************************************************************************************
// We are passing db reference connection from main to our method with other parameters
// func InsertUserTable(db *sql.DB, tablename string, user int, movie int, action int, value int) (error string) {

type UserFriendJson struct {
	Username   string `json:"username"`
	Friendname string `json:"friendname"`
}

func InsertFriendTable(c *gin.Context) {
	var msg string

	userFriendJson := UserFriendJson{}
	err := c.ShouldBindJSON(&userFriendJson)
	if err != nil {
		//fmt.Println(userFriendJson.Username + "0" + userFriendJson.Friendname)
		//fmt.Println(err.Error() + "incorrect parameters")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters, password should be between 8 to 20 chars",
		})
		return
	}

	if isFriend(userFriendJson.Username, userFriendJson.Friendname) {
		//fmt.Println(err.Error() + "nota a friend")
		c.JSON(http.StatusBadRequest, gin.H{
			"Issue": "Already an friend",
		})
	}

	user, _ := GetUserID(userFriendJson.Username)
	friend, err := GetUserID(userFriendJson.Friendname)
	if err != nil {
		//fmt.Println(err.Error() + "friendname ? ")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	//log.Println("Inserting student record ...")
	insertFriendSQL := `INSERT INTO ` + tablename3 + `(idUser, idFriend ) VALUES (?, ?)`
	statement, err := db.Prepare(insertFriendSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		//fmt.Println(err.Error() + "sql injexction ")

		//fmt.Println(err.Error())
	}

	_, err = statement.Exec(user, friend)
	if err != nil {
		//fmt.Println(err.Error() + "table entye")

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
	// Id   int
	Name string
}

var friendNames FriendNames

func GetFriends(c *gin.Context) {

	idUser, _ := GetUserID(c.Param("username"))
	// //fmt.Println(idUser)
	// //fmt.Println("disp")
	row, err := db.Query("SELECT * FROM " + tablename3 + " WHERE idUser= " + strconv.Itoa(idUser))
	if err != nil {
		//fmt.Println(err)
	}
	defer row.Close()
	var id int
	var user int
	var friend int
	var friends []FriendNames
	for row.Next() { // Iterate and fetch the records from result cursor

		row.Scan(&id, &user, &friend)
		//log.Println("Users: ", id, " ", user, " ", friend)
		// friendNames.Id = friend
		friendNames.Name = GetUserName(friend)
		//fmt.Println("sssss")
		friends = append(friends, friendNames)

	}
	//fmt.Println(friends)
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"user":    user,
		"friends": friends,
	})
}

//************************************************************************************************************************************************************************
func DisplayFriendTable(c *gin.Context) {
	//fmt.Println("disp")
	row, err := db.Query("SELECT * FROM " + tablename3 + " ORDER BY idUser")
	if err != nil {
		//fmt.Println(err)
	}
	defer row.Close()
	var id int
	var user int
	var friend int
	var friends []FriendNames
	for row.Next() { // Iterate and fetch the records from result cursor

		row.Scan(&id, &user, &friend)
		//log.Println("Users: ", id, " ", user, " ", friend)
		friendNames.Name = GetUserName(friend)
		//fmt.Println(friendNames.Name)
		friends = append(friends, friendNames)
	}
	//fmt.Println(friends)
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"user":    user,
		"friends": friends,
	})
}

//************************************************************************************************************************************************************************

func isFriend(username string, friend string) bool {
	// //fmt.Println(idUser)
	// //fmt.Println("disp")
	idUser, _ := GetUserID(username)
	idFriend, _ := GetUserID(friend)
	row, err := db.Query("SELECT COUNT(*) FROM " + tablename3 + " WHERE idUser= " + strconv.Itoa(idUser) + " AND idFriend= " + strconv.Itoa(idFriend))
	if err != nil {
		//fmt.Println(err)
	}
	defer row.Close()
	var count int
	for row.Next() { // Iterate and fetch the records from result cursor

		row.Scan(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

//************************************************************************************************************************************************************************

func GetUserList(c *gin.Context) {

	username := c.Param("username")
	// //fmt.Println(idUser)
	// //fmt.Println("disp")
	row, err := db.Query("SELECT * FROM " + tablename1 + " WHERE username LIKE '" + username + "%'")
	if err != nil {
		//fmt.Println(err)
	}
	defer row.Close()
	var id int
	var user string
	var pass string
	var friends []FriendNames
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&id, &user, &pass)
		//log.Println("Users: ", id, " ", user, " ", pass)
		// friendNames.Id = friend
		friendNames.Name = user
		// //fmt.Println("sssss")
		friends = append(friends, friendNames)

	}
	// //fmt.Println(friends)
	c.JSON(http.StatusOK, gin.H{
		"usernames": friends,
	})
}
