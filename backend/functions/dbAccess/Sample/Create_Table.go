package functions

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDB() {
	os.Remove("user_data2-database.db") // I delete the file to avoid duplicated records.
	// SQLite is a file based database.

	// log.Println("Creating sqlite-database.db...")
	file, err := os.Create("user_data-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	// log.Println("user_data-database created")

	log.Println("sqlite-database.db created")

	// Create Database Tables

	// DISPLAY INSERTED RECORDS

}

func CreateTable(db *sql.DB, tablename string) {
	createStudentTableSQL := `CREATE TABLE ` + tablename + ` (
		"idStudent" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"code" TEXT,
		"name" TEXT,
		"program" TEXT		
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
func InsertStudent(db *sql.DB, tablename string, code string, name string, program string) {
	log.Println("Inserting student record ...")
	insertStudentSQL := `INSERT INTO ` + tablename + `(code, name, program) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(code, name, program)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func DisplayStudents(db *sql.DB, tablename string) {
	row, err := db.Query("SELECT * FROM " + tablename + " ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var code string
		var name string
		var program string
		row.Scan(&id, &code, &name, &program)
		log.Println("Student: ", code, " ", name, " ", program)
	}
}
