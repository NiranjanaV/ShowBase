package tables

import (
	D "main/dbSQLite"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var tablename4 = D.GetTable(4)

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		//fmt.Println("1 Error loading .env file" + err.Error())
	}
}

type ClubShare struct {
	ClubId   int
	Password string
	Genre    int
}

//************************************************************************************************************************************************************************
func CreateClubTable() {
	createUserTableSQL := `CREATE TABLE ` + tablename4 + ` (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"password" TEXT NOT NULL,
		"genre" integer NOT NULL
	  );` // SQL Statement for Create User Table

	//log.Println("Create Share table " + tablename4)
	statement, err := db.Prepare(createUserTableSQL) // Prepare SQL Statement
	if err != nil {
		//fmt.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	//log.Println("share show table created")
}

//************************************************************************************************************************************************************************
// We are passing db reference connection from main to our method with other parameters
// func InsertUserTable(db *sql.DB, tablename string, user int, movie int, action int, value int) (error string) {
func InsertClubTable(c *gin.Context) {
	// var msg string
	type ClubShareJson struct {
		Password string `json:"password"`
		Genre    int    `json:"Genre"`
	}
	showShareJson := ClubShareJson{}
	err := c.ShouldBindJSON(&showShareJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters, password should be between 8 to 20 chars",
		})
		return
	}

	//log.Println("Inserting record ...")
	insertShareSQL := `INSERT INTO ` + tablename4 + `(password, genre ) VALUES (?, ?)`
	statement, err := db.Prepare(insertShareSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		//fmt.Println(err.Error())
	}

	_, err = statement.Exec(showShareJson.Password, showShareJson.Genre)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Your ClubID is": getClubID(showShareJson.Password, showShareJson.Genre),
	})
}

//************************************************************************************************************************************************************************

// type FriendNames struct {
// 	name string
// }

// var friendNames FriendNames

func GetClubsGenre(c *gin.Context) {
	// var password string

	type ClubShare2Json struct {
		Password string `json:"password"`
		Id       int    `json:"ClubId"`
	}
	showShare2Json := ClubShare2Json{}
	err := c.ShouldBindJSON(&showShare2Json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters, password should be between 8 to 20 chars",
		})
		return
	}

	id := showShare2Json.Id
	password := showShare2Json.Password
	// //fmt.Println(idUser)
	// //fmt.Println("disp")

	c.JSON(http.StatusOK, gin.H{
		"Genre": clubGenre(id, password),
	})
}

//************************************************************************************************************************************************************************
func DisplayClubTable(c *gin.Context) {
	//fmt.Println("disp")
	row, err := db.Query("SELECT * FROM " + tablename4 + " ORDER BY id")
	if err != nil {
		//fmt.Println(err)
	}
	defer row.Close()
	var id int
	var password string
	var genre int
	var clubShare ClubShare
	var clubShares []ClubShare
	for row.Next() { // Iterate and fetch the records from result cursor

		row.Scan(&id, &password, &genre)
		//log.Println("Clubs: ", id, " ", password, " ", genre)
		clubShare.ClubId = id
		clubShare.Password = password
		clubShare.Genre = genre
		clubShares = append(clubShares, clubShare)
	}
	c.JSON(http.StatusOK, gin.H{
		"Clubs ": clubShares,
	})
}

//************************************************************************************************************************************************************************

func isClub(id int, password string) bool {
	// //fmt.Println(idUser)
	// //fmt.Println("disp")
	row, err := db.Query("SELECT COUNT(*) FROM " + tablename4 + " WHERE id= " + strconv.Itoa(id) + " AND password= " + password)
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

func clubGenre(id int, password string) int {
	// //fmt.Println(idUser)
	// //fmt.Println("disp")
	var genre int
	if isClub(id, password) {
		row, err := db.Query("SELECT genre FROM " + tablename4 + " WHERE id= " + strconv.Itoa(id) + " AND password= " + password)
		if err != nil {
			//fmt.Println(err)
		}
		defer row.Close()

		for row.Next() { // Iterate and fetch the records from result cursor

			row.Scan(&genre)
		}

	}
	return genre
}

func getClubID(password string, genre int) int {
	// //fmt.Println(idUser)
	// //fmt.Println("disp")
	row, err := db.Query("SELECT id FROM " + tablename4 + " WHERE password= " + password + " AND genre= " + strconv.Itoa(genre))
	if err != nil {
		//fmt.Println(err)
	}
	defer row.Close()
	var id int
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&id)
	}
	return id
}
