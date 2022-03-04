package tables

import (
	"database/sql"
	"fmt"
	"log"
	D "main/dbSQLite"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var db = D.GetDB()
var tablename = D.GetTable(1)

func CreateAuthTable(db *sql.DB, tablename string) {
	createStudentTableSQL := `CREATE TABLE ` + tablename + ` (
		"idUser" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username" TEXT UNIQUE,
		"password" TEXT		
	  );` // SQL Statement for Create Table

	log.Println("Create student table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("student table created")
}

// We are passing db reference connection from main to our method with other parameters
func InsertAuthTable(c *gin.Context) {
	var msg string
	type UserAuthJson struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	userAuthJson := UserAuthJson{}
	err := c.ShouldBindJSON(&userAuthJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters, password should be between 8 to 20 chars",
		})
		return
	}

	log.Println("Inserting student record ...")
	insertStudentSQL := `INSERT INTO ` + tablename + `(username, password) VALUES (?, ?)`

	statement, err := db.Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	_, err = statement.Exec(userAuthJson.Username, userAuthJson.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		msg = "inserted"
	}
	c.JSON(http.StatusOK, gin.H{
		"Return": msg,
	})
}

func GetPassForUser(c *gin.Context) {
	var auth int
	type UserAuthJson struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	userAuthJson := UserAuthJson{}
	err := c.ShouldBindJSON(&userAuthJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters, password should be between 8 to 20 chars",
		})
		return
	}

	fmt.Println("get")
	row, err := db.Query("SELECT * FROM " + tablename + " WHERE USERNAME = '" + userAuthJson.Username + "'")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var username string
		var password []byte
		row.Scan(&id, &username, &password)
		log.Println("Users: ", id, " ", username, " ", password)
		if err := bcrypt.CompareHashAndPassword(password, []byte(userAuthJson.Password)); err != nil {
			// TODO: Properly handle error
			log.Fatal(err)
			auth = 0
		} else {
			auth = 1
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"Return": auth,
	})
}

func DisplayAuthTable(c *gin.Context) {
	fmt.Println("disp")
	row, err := db.Query("SELECT * FROM " + tablename + " ORDER BY username")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var username string
		var password string
		row.Scan(&id, &username, &password)
		log.Println("Users: ", id, " ", username, " ", password)
		c.JSON(http.StatusOK, gin.H{
			"ID":       id,
			"Username": username,
			"Password": password,
		})
	}
}
