package data

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

//chnage movie struct if needed ,,,its different for queries
// type FilmId struct {
// 	Poster_path       string
// 	Adult             bool
// 	Overview          string
// 	Release_date      string
// 	Genres            []Genre
// 	Id                int32
// 	Original_title    string
// 	Original_language string
// 	Vote_average      float32
// 	Title             string
// 	Backdrop_path     string
// 	Popularity        float64
// 	Vote_count        int32
// 	Video             bool
// }

// type Genre struct {
// 	Id   int32
// 	Name string
// }

// // var api Api
// var filmId FilmId
// var genr []Genre

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

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
func InsertUserTable(db *sql.DB, tablename string, user int, movie int, action int, value int) (error string) {
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
			// gen := []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

			//Code to get genre *****************************************************************************************************************************

			gen := GetGenre(movie)
			// response, err := http.Get("https://api.themoviedb.org/3/movie/" + strconv.Itoa(movie) + "?api_key=" + os.Getenv("Twilio_api_key") + "&language=en-US")

			// if err != nil {
			// 	fmt.Print(err.Error())
			// 	os.Exit(1)
			// }
			// // fmt.Println(response)

			// responseData, err := ioutil.ReadAll(response.Body)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// // fmt.Println(responseData)

			// json.Unmarshal([]byte(responseData), &filmId)
			// fmt.Println(filmId)

			// // json.Unmarshal([]byte(filmId.Genres), &genr)
			// // fmt.Println(genr)

			// data, _ := json.Marshal(filmId.Genres)

			// json.Unmarshal([]byte(data), &genr)
			// fmt.Println(genr)

			// // for genres := range genr {
			// // 	ids, _ := json.Marshal(genres.Id)
			// // 	fmt.Println(string(ids))
			// // }

			// for i, s := range genr {
			// 	var g int
			// 	genres, _ := json.Marshal(s.Id)
			// 	json.Unmarshal([]byte(genres), &g)
			// 	gen[i] = g

			// }
			fmt.Println(gen)
			//Code to get genre end *************************************************************************************************************************

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

			_, err = statement.Exec(value, user, movie, gen[0], gen[1], gen[2])
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
