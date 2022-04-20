package tables

import (
	D "main/dbSQLite"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

//************************************************************************************************************************************************************************

var db = D.GetDB()
var tablename = D.GetTable(1)

func CreateAuthTable() {
	createStudentTableSQL := `CREATE TABLE ` + tablename + ` (
		"idUser" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username" TEXT UNIQUE,
		"password" TEXT		
	  );`

	// SQL Statement for Create Table

	//log.Println("Create student table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		//fmt.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	//log.Println("student table created")
}

//************************************************************************************************************************************************************************

// We are passing db reference connection from main to our method with other parameters
func InsertAuthTable(c *gin.Context) {
	var msg string
	var db = D.GetDB()
	var tablename = D.GetTable(1)
	// //fmt.Println("aaaaaaaaaaaaaaaa", db.Ping())

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

	//log.Println("Inserting student record ...")

	//checking if user exists
	row, err := db.Query("SELECT COUNT(*) FROM " + tablename + " WHERE USERNAME = '" + userAuthJson.Username + "'")
	if err != nil {
		//fmt.Print("User check issue exists")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	for row.Next() { // Iterate and fetch the records from result cursor
		var count int
		row.Scan(&count)
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "User already exists, try another",
			})
		} else {
			row.Close()
			//fmt.Print("Insertion Step")

			insertStudentSQL := `INSERT INTO ` + tablename + ` (username, password) VALUES (?, ?)`

			statement, err := db.Prepare(insertStudentSQL) // Prepare statement.
			// This is good to avoid SQL injections
			if err != nil {
				//fmt.Println(err.Error())
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}

			hash, _ := bcrypt.GenerateFromPassword([]byte(userAuthJson.Password), bcrypt.DefaultCost)
			_, err = statement.Exec(userAuthJson.Username, hash)
			if err != nil {
				//fmt.Print("Error inserting", err.Error())

				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				msg = "inserted"
				//fmt.Print("inserted")

			}
			c.JSON(http.StatusOK, gin.H{
				"Return": msg,
			})

		}

	}

}

//************************************************************************************************************************************************************************

func GetPassForUser(c *gin.Context) {
	var auth int
	//fmt.Println("get1")

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

	//fmt.Println("get2")
	row, err := db.Query("SELECT * FROM " + tablename + " WHERE USERNAME = '" + userAuthJson.Username + "'")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	//fmt.Println("get3")
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var username string
		var password []byte
		row.Scan(&id, &username, &password)
		//log.Println("Users: ", id, " ", username, " ", password)
		//fmt.Println(password)
		if err := bcrypt.CompareHashAndPassword(password, []byte(userAuthJson.Password)); err != nil {
			// TODO: Properly handle error
			//fmt.Println(err)
			auth = 0
		} else {
			auth = 1
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"Return": auth,
	})
}

//************************************************************************************************************************************************************************

func DisplayAuthTable(c *gin.Context) {
	//fmt.Println("disp")
	row, err := db.Query("SELECT * FROM " + tablename + " ORDER BY username")
	if err != nil {
		//fmt.Println(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var username string
		var password string
		row.Scan(&id, &username, &password)
		//log.Println("Users: ", id, " ", username, " ", password)
		c.JSON(http.StatusOK, gin.H{
			"ID":       id,
			"Username": username,
			"Password": password,
		})
	}
}

//************************************************************************************************************************************************************************

func InsertAuthTableGo(username string, password string) {

	insertStudentSQL := `INSERT INTO ` + tablename + `(username, password) VALUES (?, ?)`

	statement, err := db.Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		//fmt.Println("DB issue")
		//fmt.Println("error", err.Error())
	}
	_, err = statement.Exec(username, password)
	if err != nil {
		//fmt.Println("error", err.Error())

	} else {
		// msg := "inserted"
		//fmt.Println(msg)

	}
}

//************************************************************************************************************************************************************************

func DisplayAuthTableGo() {
	//fmt.Println("disp")
	row, err := db.Query("SELECT * FROM " + tablename)
	if err != nil {
		//fmt.Println(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var username string
		var password string
		row.Scan(&id, &username, &password)
		//log.Println("Users: ", id, " ", username, " ", password)

	}
}

//************************************************************************************************************************************************************************

func PingDBGo() {
	//fmt.Println(db.Ping())
}
