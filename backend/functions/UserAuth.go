package functions

import (
	// M "/GoCode/Create_Table.go"
	// N "/GoCode/controllers"
	// L "ShowBase/data"
	"database/sql"
	"fmt"

	//O "functions/databaseEntry"

	"log"
	// "os"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func main2() {
	// var database *sql.DB
	// os.Remove("user_data-database.db")                               // I delete the file to avoid duplicated records.
	sqliteDatabase, _ := sql.Open("sqlite3", "./../dbSQLite/user_data-database") // Open the created SQLite File
	defer sqliteDatabase.Close()
	// database = sqliteDatabase
	var tablename = "user_auth"

	//to get from frontend
	var username = "username2"
	var password = "password2"

	// Create DB
	// fmt.Println("Create db")
	// O.CreateDB()

	// Create Table
	// fmt.Println("Create table")
	// O.CreateAuthTable(sqliteDatabase, tablename) // Create Database Tables

	// INSERT RECORDS
	fmt.Println("insert")
	//hashing pass
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	fmt.Println("Hash to store:", string(hash))
	error := InsertAuthTable(sqliteDatabase, tablename, username, hash)
	if error != "None" {
		fmt.Println("Duplicate username")
	}

	// Authenticate user
	newPass := "password"
	auth := GetPassForUser(sqliteDatabase, tablename, username, newPass)
	if auth == 0 {
		fmt.Print("User not authenticated")
	} else {
		fmt.Println("User Authenticated")
	}

	// DISPLAY INSERTED RECORDS
	fmt.Println("display")
	DisplayAuthTable(sqliteDatabase, tablename)
}
