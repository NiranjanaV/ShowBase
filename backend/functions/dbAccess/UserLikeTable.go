package tables

import (
	"encoding/json"
	"io/ioutil"
	D "main/dbSQLite"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var tablename2 = D.GetTable(2)

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		//fmt.Println("1 Error loading .env file" + err.Error())
	}
}

//************************************************************************************************************************************************************************
func CreateUserTable() {
	createUserTableSQL := `CREATE TABLE user3 (
		"idLike" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"idUser" integer NOT NULL,
		"movieID" integer NOT NULL,
		"like" integer DEFAULT -1,
		"watched" integer DEFAULT -1,
		"watching" integer DEFAULT -1,
		"toWatch" integer DEFAULT -1,
		"genre1" integer DEFAULT -1,
		"genre2" integer DEFAULT -1,
		"genre3" integer DEFAULT -1,
		FOREIGN KEY (idUser) REFERENCES user_auth(idUser),
		CHECK (watching>=-1 AND watching <=100 AND
			like>=-1 AND like <=10 AND
			watched IN (-1,0,1,NULL) AND
			toWatch IN (-1,0,1,NULL) 
		)

	  );` // SQL Statement for Create User Table

	//log.Println("Create user table...")
	statement, err := db.Prepare(createUserTableSQL) // Prepare SQL Statement
	if err != nil {
		//fmt.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	//log.Println("user table created")
}

//************************************************************************************************************************************************************************

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

var filmId FilmId

var genr []Genre

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		//fmt.Println("6 Error loading .env file" + err.Error())
	}
}

func GetGenre2(movieID int) (gen []int) {

	response, err := http.Get("https://api.themoviedb.org/3/movie/" + strconv.Itoa(movieID) + "?api_key=" + os.Getenv("Twilio_api_key") + "&language=en-US")

	if err != nil {
		//fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//fmt.Println(err)
	}

	json.Unmarshal([]byte(responseData), &filmId)

	data, _ := json.Marshal(filmId.Genres)

	json.Unmarshal([]byte(data), &genr)

	gen = []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	for i, s := range genr {
		var g int
		genres, _ := json.Marshal(s.Id)
		json.Unmarshal([]byte(genres), &g)
		gen[i] = g
	}
	// //fmt.Println(gen)

	return
}

//************************************************************************************************************************************************************************
// We are passing db reference connection from main to our method with other parameters
// func InsertUserTable(db *sql.DB, tablename string, user int, movie int, action int, value int) (error string) {
func InsertUserTable(c *gin.Context) {
	//fmt.Printf("    in         dd")
	var msg string
	type UserLikeJson struct {
		Username string `json:"username"`
		Movie    int    `json:"movie"`
		Action   int    `json:"action"`
		Value    int    `json:"value"`
	}
	userLikeJson := UserLikeJson{}
	err := c.ShouldBindJSON(&userLikeJson)
	if err != nil {
		//fmt.Println("error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "fffffffffffffffffffff",
		})
		return
	}

	user, _ := GetUserID(userLikeJson.Username)
	movie := userLikeJson.Movie
	action := userLikeJson.Action
	value := userLikeJson.Value

	//fmt.Println(userLikeJson.Username, strconv.Itoa(movie), strconv.Itoa(action), strconv.Itoa(value))
	// //fmt.Println(userLikeJson.Username)
	//log.Println("Inserting student record ...")
	insertUserSQL := `INSERT INTO ` + tablename2 + `(idUser, movieID, like ) VALUES (?, ?, ?)`

	//fmt.Println("get")
	row, err := db.Query("SELECT COUNT(*) FROM " + tablename2 + " WHERE idUser = " + strconv.Itoa(user) + " AND movieID = " + strconv.Itoa(movie))
	if err != nil {
		//fmt.Println(err)
	}
	//defer
	for row.Next() { // Iterate and fetch the records from result cursor
		var count int
		row.Scan(&count)
		row.Close()
		//log.Println("Users: ", count)
		if count > 0 {
			// TODO: Properly handle error
			// //fmt.Println(err)
			// auth = 0

			if action == 1 {
				insertUserSQL = `UPDATE ` + tablename2 + ` SET like= ? WHERE  idUser = ? and movieID = ?`
			} else if action == 2 {
				insertUserSQL = `UPDATE ` + tablename2 + ` SET watched= ? WHERE  idUser = ? and movieID = ?`
			} else if action == 3 {
				insertUserSQL = `UPDATE ` + tablename2 + ` SET watching= ? WHERE  idUser = ? and movieID = ?`
			} else if action == 4 {
				insertUserSQL = `UPDATE ` + tablename2 + ` SET toWatch= ? WHERE  idUser = ? and movieID = ?`
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid action",
				})
			}
			statement, err := db.Prepare(insertUserSQL) // Prepare statement.
			// This is good to avoid SQL injections
			if err != nil {
				//fmt.Println(err.Error())
			}
			_, err = statement.Exec(value, user, movie)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				//fmt.Println("inserted")
				msg = "Inserted"
			}

		} else {
			//Code to get genre *****************************************************************************************************************************
			//fmt.Println("gengengengengengengengengengengengengengengengengen")
			gen := GetGenre2(movie)
			//fmt.Println(gen)

			//Code to get genre end *************************************************************************************************************************

			if action == 1 {
				//fmt.Println("action as like aaaaaaaaaaaaaaaaaaaaaaaaaa")
				insertUserSQL = `INSERT INTO ` + tablename2 + `( like, idUser, movieID, genre1, genre2, genre3 ) VALUES (?, ?, ?, ?, ?, ?)`
			} else if action == 2 {
				insertUserSQL = `INSERT INTO ` + tablename2 + `( watched, idUser, movieID, genre1, genre2, genre3 ) VALUES (?, ?, ?, ?, ?, ?)`
			} else if action == 3 {
				insertUserSQL = `INSERT INTO ` + tablename2 + `( watching, idUser, movieID,  genre1, genre2, genre3 ) VALUES (?, ?, ?, ?, ?, ?)`
			} else if action == 4 {
				insertUserSQL = `INSERT INTO ` + tablename2 + `( toWatch, idUser, movieID,  genre1, genre2, genre3 ) VALUES (?, ?, ?, ?, ?, ?)`
			} else {
				//fmt.Println("Insert command", err.Error())
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid Action",
				})
			}
			// //fmt.Println(db.Prepare(insertUserSQL))
			statement, err := db.Prepare(insertUserSQL) // Prepare statement.
			// This is good to avoid SQL injections
			if err != nil {
				//fmt.Println("SQL Injection: ", err.Error())
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}

			_, err = statement.Exec(value, user, movie, gen[0], gen[1], gen[2])
			//fmt.Println(sqlresult)
			if err != nil {
				//fmt.Println("Executing statement : ", err.Error())
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				//fmt.Println("inserted")
				msg = "Inserted"
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"Return": msg,
	})
}

//************************************************************************************************************************************************************************

func GetUserTable(c *gin.Context) {

	idUser, _ := GetUserID(c.Param("username"))
	//fmt.Println(idUser)
	//fmt.Println("disp")
	row, err := db.Query("SELECT * FROM " + tablename2 + " WHERE idUser= " + strconv.Itoa(idUser))
	if err != nil {
		//fmt.Println(err)
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
		//log.Println("Users: ", id, " ", user, " ", movie, " ", like, " ", watched, " ", watching, " ", toWatch, " ", genre1, " ", genre2, " ", genre3)
		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"movie":    movie,
			"user":     user,
			"like":     like,
			"watched":  watched,
			"watching": watching,
			"toWatch":  toWatch,
			"genre1":   genre1,
			"genre2":   genre2,
			"genre3":   genre3,
		})
	}
}

//************************************************************************************************************************************************************************
func DisplayUserTable(c *gin.Context) {
	//fmt.Println("disp")
	row, err := db.Query("SELECT * FROM " + tablename2 + " ORDER BY idUser")
	if err != nil {
		//fmt.Println(err)
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
		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"movie":    movie,
			"user":     user,
			"like":     like,
			"watched":  watched,
			"watching": watching,
			"toWatch":  toWatch,
			"genre1":   genre1,
			"genre2":   genre2,
			"genre3":   genre3,
		})
		//log.Println("Users: ", id, " ", user, " ", movie, " ", like, " ", watched, " ", watching, " ", toWatch, " ", genre1, " ", genre2, " ", genre3)
	}
}

//************************************************************************************************************************************************************************

type UserPref struct {
	Id       int
	User     int
	Movie    int
	Like     int
	Watched  int
	Watching int
	ToWatch  int
	Genre1   int
	Genre2   int
	Genre3   int
}

var userPref UserPref

func DisplayUsersMovie(user int, movie int) UserPref {
	//fmt.Println("get user and his movie")
	row, err := db.Query("SELECT COUNT(*) FROM " + tablename2 + " WHERE idUser='" + strconv.Itoa(user) + "' AND movieID='" + strconv.Itoa(movie) + "'ORDER BY idUser")
	if err != nil {
		//fmt.Println(err)
	}
	defer row.Close()

	// var slice = make([]int, 0, 10)
	var count int
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&count)
		if count == 1 {
			row, err := db.Query("SELECT * FROM " + tablename2 + " WHERE idUser='" + strconv.Itoa(user) + "' AND movieID='" + strconv.Itoa(movie) + "'ORDER BY idUser")
			if err != nil {
				//fmt.Println(err)
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
				userPref.Id = id
				userPref.User = user
				userPref.Movie = movie
				userPref.Like = like
				userPref.Watched = watched
				userPref.Watching = watching
				userPref.ToWatch = toWatch
				userPref.Genre1 = genre1
				userPref.Genre2 = genre2
				userPref.Genre3 = genre3
			}
		}
	}
	return userPref
}
