package main

import (
	// M "/GoCode/Create_Table.go"
	// N "/GoCode/controllers"
	// L "ShowBase/data"
	"database/sql"
	"fmt"
	O "main/data"

	// "log"
	// "os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// var database *sql.DB
	// os.Remove("user_data-database.db")                               // I delete the file to avoid duplicated records.
	sqliteDatabase, _ := sql.Open("sqlite3", "./user_data-database") // Open the created SQLite File
	defer sqliteDatabase.Close()
	// database = sqliteDatabase
	var tablename = "user_auth"

	// fmt.Println("Create db")
	// O.CreateDB()

	fmt.Println("Create table")
	O.CreateTable(sqliteDatabase, tablename) // Create Database Tables

	// INSERT RECORDS
	fmt.Println("insert")
	O.InsertStudent(sqliteDatabase, tablename, "1", "username", "password")

	// DISPLAY INSERTED RECORDS
	fmt.Println("display")
	O.DisplayStudents(sqliteDatabase, tablename)
}
