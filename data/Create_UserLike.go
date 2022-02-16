package data

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

//************************************************************************************************************************************************************************
func CreateUserTable(db *sql.DB, tablename string) {
	createUserTableSQL := `CREATE TABLE ` + tablename + ` (
		"idLike" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"idUser" integer NOT NULL,
		"movieID" integer NOT NULL,
		"like" integer DEFAULT NULL,
		"watched" integer DEFAULT NULL,
		"watching" integer DEFAULT NULL,
		"toWatch" integer DEFAULT NULL,
		"genre1" integer DEFAULT NULL,
		"genre2" integer DEFAULT NULL,
		"genre3" integer DEFAULT NULL,
		FOREIGN KEY (idUser) REFERENCES user_auth(idUser),
		CHECK (watching>=0 AND watching <=100 AND
			like>=0 AND like <=10 AND
			watched IN (0,1,NULL) AND
			toWatch IN (0,1,NULL) 
		)

	  );` // SQL Statement for Create User Table

	log.Println("Create user table...")
	statement, err := db.Prepare(createUserTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("user table created")
}

//************************************************************************************************************************************************************************
// We are passing db reference connection from main to our method with other parameters
func InsertUserTable(db *sql.DB, tablename string, user int, movie int, action int, value int, gen1 int, gen2 int, gen3 int) (error string) {
	log.Println("Inserting student record ...")
	insertUserSQL := `INSERT INTO ` + tablename + `(idUser, movieID, like ) VALUES (?, ?, ?)`

	fmt.Println("get")
	row, err := db.Query("SELECT COUNT(*) FROM " + tablename + " WHERE idUser = " + strconv.Itoa(user) + " AND movieID = " + strconv.Itoa(movie))
	if err != nil {
		log.Fatal(err)
	}
	//defer
	for row.Next() { // Iterate and fetch the records from result cursor
		var count int
		row.Scan(&count)
		row.Close()
		log.Println("Users: ", count)
		if count > 0 {
			// TODO: Properly handle error
			// log.Fatal(err)
			// auth = 0

			if action == 1 {
				insertUserSQL = `UPDATE ` + tablename + ` SET like= ? WHERE  idUser = ? and movieID = ?`
			} else if action == 2 {
				insertUserSQL = `UPDATE ` + tablename + ` SET watched= ? WHERE  idUser = ? and movieID = ?`
			} else if action == 3 {
				insertUserSQL = `UPDATE ` + tablename + ` SET watching= ? WHERE  idUser = ? and movieID = ?`
			} else if action == 4 {
				insertUserSQL = `UPDATE ` + tablename + ` SET toWatch= ? WHERE  idUser = ? and movieID = ?`
			} else {
				error = "Invalid action"
			}
			statement, err := db.Prepare(insertUserSQL) // Prepare statement.
			// This is good to avoid SQL injections
			if err != nil {
				log.Fatalln(err.Error())
			}
			_, err = statement.Exec(value, user, movie)
			if err != nil {
				error = err.Error()
			} else {
				fmt.Println("inserted")
				error = "None"
			}

		} else {

			if action == 1 {
				insertUserSQL = `INSERT INTO ` + tablename + `( like, idUser, movieID, genre1, genre2, genre3 ) VALUES (?, ?, ?, ?, ?, ?)`
			} else if action == 2 {
				insertUserSQL = `INSERT INTO ` + tablename + `( watched, idUser, movieID, genre1, genre2, genre3 ) VALUES (?, ?, ?, ?, ?, ?)`
			} else if action == 3 {
				insertUserSQL = `INSERT INTO ` + tablename + `( watching, idUser, movieID,  genre1, genre2, genre3 ) VALUES (?, ?, ?, ?, ?, ?)`
			} else if action == 4 {
				insertUserSQL = `INSERT INTO ` + tablename + `( toWatch, idUser, movieID,  genre1, genre2, genre3 ) VALUES (?, ?, ?, ?, ?, ?)`
			} else {
				error = "Invalid action"
			}
			statement, err := db.Prepare(insertUserSQL) // Prepare statement.
			// This is good to avoid SQL injections
			if err != nil {
				log.Fatalln(err.Error())
			}

			_, err = statement.Exec(value, user, movie, gen1, gen2, gen3)
			if err != nil {
				error = err.Error()
			} else {
				fmt.Println("inserted")
				error = "None"
			}
		}
	}
	return
}

//************************************************************************************************************************************************************************
// func GetPassForUser(db *sql.DB, tablename string, user string, newPass string) (auth int) {
// 	fmt.Println("get")
// 	row, err := db.Query("SELECT * FROM " + tablename + " WHERE USERNAME = '" + user + "'")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer row.Close()
// 	for row.Next() { // Iterate and fetch the records from result cursor
// 		var id int
// 		var username string
// 		var password []byte
// 		row.Scan(&id, &username, &password)
// 		log.Println("Users: ", id, " ", username, " ", password)
// 		if err := bcrypt.CompareHashAndPassword(password, []byte(newPass)); err != nil {
// 			// TODO: Properly handle error
// 			log.Fatal(err)
// 			auth = 0
// 		} else {
// 			auth = 1
// 		}
// 	}
// 	return
// }

//************************************************************************************************************************************************************************
func GetUserTable(db *sql.DB, tablename string, idUser int) (readUser string) {
	fmt.Println("disp")
	row, err := db.Query("SELECT * FROM " + tablename + " WHERE idUser= " + strconv.Itoa(idUser))
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var user int
		var movie int
		var like int
		var watched int
		var watching int
		var toWatch int
		var genre1 int
		var genre2 int
		var genre3 int
		row.Scan(&id, &user, &movie, &like, &watched, &watching, &toWatch, &genre1, &genre2, &genre3)
		log.Println("Users: ", id, " ", user, " ", movie, " ", like, " ", watched, " ", watching, " ", toWatch, " ", genre1, " ", genre2, " ", genre3)
		readUser = "Users: " + strconv.Itoa(id) + " " + strconv.Itoa(user) + " " + strconv.Itoa(movie) + " " + strconv.Itoa(like) + " " + strconv.Itoa(watched) + " " + strconv.Itoa(watching) + " " + strconv.Itoa(toWatch) + " " + strconv.Itoa(genre1) + " " + strconv.Itoa(genre2) + " " + strconv.Itoa(genre3)
	}
	return
}

//************************************************************************************************************************************************************************
func DisplayUserTable(db *sql.DB, tablename string) {
	fmt.Println("disp")
	row, err := db.Query("SELECT * FROM " + tablename + " ORDER BY idUser")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var user int
		var movie int
		var like int
		var watched int
		var watching int
		var toWatch int
		var genre1 int
		var genre2 int
		var genre3 int
		row.Scan(&id, &user, &movie, &like, &watched, &watching, &toWatch, &genre1, &genre2, &genre3)
		log.Println("Users: ", id, " ", user, " ", movie, " ", like, " ", watched, " ", watching, " ", toWatch, " ", genre1, " ", genre2, " ", genre3)
	}
}

//************************************************************************************************************************************************************************
