package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

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
func InsertAuthTable(db *sql.DB, tablename string, user string, password []byte) (error string) {
	log.Println("Inserting student record ...")
	insertStudentSQL := `INSERT INTO ` + tablename + `(username, password) VALUES (?, ?)`

	statement, err := db.Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(user, password)
	if err != nil {
		error = err.Error()
	} else {
		fmt.Println("inserted")
		error = "None"
	}
	return
}

func GetPassForUser(db *sql.DB, tablename string, user string, newPass string) (auth int) {

	fmt.Println("get")
	row, err := db.Query("SELECT * FROM " + tablename + " WHERE USERNAME = '" + user + "'")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var username string
		var password []byte
		row.Scan(&id, &username, &password)
		log.Println("Users: ", id, " ", username, " ", password)
		if err := bcrypt.CompareHashAndPassword(password, []byte(newPass)); err != nil {
			// TODO: Properly handle error
			log.Fatal(err)
			auth = 0
		} else {
			auth = 1
		}
	}
	return
}

func DisplayAuthTable(db *sql.DB, tablename string) {
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
	}
}
