package tables

import (
	D "main/dbSQLite"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var tablename5 = D.GetTable(5)

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		//fmt.Println("1 Error loading .env file" + err.Error())
	}
}

type Club struct {
	ClubId int
	Movies []MovieShare
}

type MovieShare struct {
	Movie int
	Value int
}

//************************************************************************************************************************************************************************
func CreateMovieTable() {
	createUserTableSQL := `CREATE TABLE ` + tablename5 + ` (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"club" integer NOT NULL,
		"movie" integer NOT NULL,
		"vote" integer DEFAULT 0,
		"value" integer DEFAULT 0,
		FOREIGN KEY(culb) REFERENCES share_club(id)
	  );` // SQL Statement for Create User Table

	//log.Println("Create movie table " + tablename4)
	statement, err := db.Prepare(createUserTableSQL) // Prepare SQL Statement
	if err != nil {
		//fmt.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	//log.Println("share show table created")
}

//************************************************************************************************************************************************************************
// We are passing db reference connection from main to our method with other parameters
// func InsertUserTable(db *sql.DB, tablename string, user int, movie int, action int, value int) (error string) {
func InsertMovieTable(c *gin.Context) {
	// var msg string
	type MovieVoteJson struct {
		Club  int `json:"club"`
		Movie int `json:"movie"`
		Value int `json:"value"`
	}
	showShareJson := []MovieVoteJson{}
	err := c.ShouldBindJSON(&showShareJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters, password should be between 8 to 20 chars",
		})
		return
	}
	for i := 0; i < len(showShareJson); i++ {
		if moviePresent(showShareJson[i].Club, showShareJson[i].Movie) {
			//log.Println("Updating records ...")
			vote, value := getClubMovie(showShareJson[i].Club, showShareJson[i].Movie)
			updateMovieSQL := `UPDATE ` + tablename5 + ` SET value = ?, vote = ? WHERE club= ` + strconv.Itoa(showShareJson[i].Club) + ` AND movie= ` + strconv.Itoa(showShareJson[i].Movie) + ``
			statement, err := db.Prepare(updateMovieSQL) // Prepare statement.
			// This is good to avoid SQL injections
			if err != nil {
				//fmt.Println(err.Error())
			}
			_, err = statement.Exec(value+showShareJson[i].Value, vote+1)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}
		} else {
			//log.Println("Inserting record ...")
			insertShareSQL := `INSERT INTO ` + tablename5 + `(club, movie, value, vote ) VALUES (? , ?, ?, ?)`
			statement, err := db.Prepare(insertShareSQL) // Prepare statement.
			// This is good to avoid SQL injections
			if err != nil {
				//fmt.Println(err.Error())
			}
			_, err = statement.Exec(showShareJson[i].Club, showShareJson[i].Movie, showShareJson[i].Value, 1)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": "Inserted",
	})
}

//************************************************************************************************************************************************************************

// type FriendNames struct {
// 	name string
// }

// var friendNames FriendNames

// func GetClubsGenre(c *gin.Context) {
// 	// var password string

// 	type ClubShare2Json struct {
// 		Password string `json:"password"`
// 		Id       int    `json:"ClubId"`
// 	}
// 	showShare2Json := ClubShare2Json{}
// 	err := c.ShouldBindJSON(&showShare2Json)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "incorrect parameters, password should be between 8 to 20 chars",
// 		})
// 		return
// 	}

// 	id := showShare2Json.Id
// 	password := showShare2Json.Password
// 	// //fmt.Println(idUser)
// 	// //fmt.Println("disp")

// 	c.JSON(http.StatusOK, gin.H{
// 		"Genre": clubGenre(id, password),
// 	})
// }

//************************************************************************************************************************************************************************
func DisplayMovieTable(c *gin.Context) {
	//fmt.Println("disp")
	row, err := db.Query("SELECT * FROM " + tablename5 + " ORDER BY id")
	if err != nil {
		//fmt.Println(err)
	}
	defer row.Close()

	var clubId int
	var movie int
	var vote int
	var value int
	var avgValue int

	var clubSingle Club
	var movieShare MovieShare
	var clubs []Club
	clubMoviePairs := make(map[int][]MovieShare)
	for row.Next() { // Iterate and fetch the records from result cursor

		row.Scan(&clubId, &movie, &vote, &value)
		avgValue = value / vote
		//log.Println("Club: ", clubId, " ", movie, " ", avgValue)
		movieShare.Movie = movie
		movieShare.Value = avgValue
		movies := clubMoviePairs[clubId]
		movies = append(movies, movieShare)
		clubMoviePairs[clubId] = movies
	}
	var i = 0
	for club, movieShare := range clubMoviePairs {
		clubSingle.ClubId = club
		clubSingle.Movies = movieShare
		i++
		clubs = append(clubs, clubSingle)
	}

	c.JSON(http.StatusOK, gin.H{
		"Clubs ": clubs,
	})
}

//************************************************************************************************************************************************************************

func moviePresent(club int, movie int) bool {
	// //fmt.Println(idUser)
	// //fmt.Println("disp")
	row, err := db.Query("SELECT COUNT(*) FROM " + tablename5 + " WHERE club= " + strconv.Itoa(club) + " AND movie= " + strconv.Itoa(movie))
	if err != nil {
		//fmt.Println(err)
	}
	defer row.Close()
	var count int
	for row.Next() { // Iterate and fetch the records from result cursor

		row.Scan(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

func getClubMovie(club int, movie int) (vote int, value int) {
	// //fmt.Println(idUser)
	// //fmt.Println("disp")
	row, err := db.Query("SELECT vote, value FROM " + tablename5 + " WHERE club= " + strconv.Itoa(club) + " AND movie= " + strconv.Itoa(movie))
	if err != nil {
		//fmt.Println(err)
	}
	defer row.Close()

	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&vote, &value)
	}
	return
}
