package main

import (
	"database/sql"
	"log"
	"os"
	 L "./\C:\Users\ksrik\Documents\GitHub\ShowBase"
	_ "github.com/mattn/go-sqlite3"
)


func main (){
	var database *sql.DB
	sqliteDatabase, _ := sql.Open("sqlite3", "./user_data-database") // Open the created SQLite File
	defer sqliteDatabase.Close()
	database =sqliteDatabase
	var tablename = "user_preference" 
	L.createTable(sqliteDatabase, tablename) // Create Database Tables

	// INSERT RECORDS
	L.insertStudent(database, "0001", "Liana Kim", "Bachelor")
	

	// DISPLAY INSERTED RECORDS
	L.displayStudents(sqliteDatabase)
}