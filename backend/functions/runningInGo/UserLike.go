package functions

import (

	//O "main/data"

	// "os"

	_ "github.com/mattn/go-sqlite3"
)

// func main3() {
// 	// var database *sql.DB
// 	// os.Remove("user_data-database.db")                               // I delete the file to avoid duplicated records.
// 	sqliteDatabase, _ := sql.Open("sqlite3", "./../dbSQLite/user_data-database") // Open the created SQLite File
// 	defer sqliteDatabase.Close()
// 	// database = sqliteDatabase
// 	var tablename = "user"

// 	//to get from frontend
// 	user := 1
// 	movie := 353044
// 	like := 1
// 	// watched := 1
// 	// watching := 1
// 	// toWatch := 1

// 	// Create DB****************************************************************************************
// 	// fmt.Println("Create db")
// 	// O.CreateDB()

// 	// Create Table ************************************************************************************

// 	// fmt.Println("Create table")
// 	// O.CreateUserTable(sqliteDatabase, tablename) // Create Database Tables

// 	// INSERT RECORDS*****************************************************************************************
// 	error := InsertUserTable(sqliteDatabase, tablename, user, movie, 1, like)
// 	if error != "None" {
// 		fmt.Println("Error inserting " + error)
// 	}

// 	error = InsertUserTable(sqliteDatabase, tablename, user, movie, 2, like)
// 	if error != "None" {
// 		fmt.Println("Error inserting " + error)
// 	}

// 	error = InsertUserTable(sqliteDatabase, tablename, 2, movie, 4, like)
// 	if error != "None" {
// 		fmt.Println("Error inserting " + error)
// 	}

// 	// get user RECORDS*****************************************************************************************
// 	fmt.Println("get user 1")
// 	userstring := GetUserTable(sqliteDatabase, tablename, user)
// 	fmt.Println("STRING IN MAIN : " + userstring)

// 	// DISPLAY INSERTED RECORDS*****************************************************************************************
// 	fmt.Println("display")
// 	DisplayUserTable(sqliteDatabase, tablename)
// }
