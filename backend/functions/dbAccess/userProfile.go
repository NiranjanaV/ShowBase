package tables

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	// D "main/dbSQLite"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

type Genre struct {
	Id   int32
	Name string
}

type FilmId struct {
	Poster_path       string
	Adult             bool
	Overview          string
	Release_date      string
	Genres            []Genre
	Id                int32
	Original_title    string
	Original_language string
	Vote_average      float32
	Title             string
	Backdrop_path     string
	Popularity        float64
	Vote_count        int32
	Video             bool
}

type UserLike struct {
	userLikeVal   int
	userLikeMovie FilmId
}

type UserWatching struct {
	userWatchingVal   int
	userWatchingMovie FilmId
}

//************************************************************************************************************************************************************************

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		log.Fatal("1 Error loading .env file" + err.Error())
	}
}

//************************************************************************************************************************************************************************

func SendUserProfile(c *gin.Context) {

	var userLikeSingle UserLike
	var userLike []UserLike
	var userWatched []FilmId
	var userWatingSingle UserWatching
	var userWatching []UserWatching
	var userToWatch []FilmId

	username := c.Param("username")
	idUser := GetUserID(c.Param("username"))
	fmt.Println(idUser)
	fmt.Println("disp")
	//************************************************************************************************************************************************************************

	//watching
	row, err := db.Query("SELECT movieID, watching FROM " + tablename2 + " WHERE idUser= " + strconv.Itoa(idUser) + " AND watching <> -1")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		// var id int
		// var user int
		var movie int
		// var like int
		// var watched int
		var watching int
		// var toWatch int
		// var genre1 int
		// var genre2 int
		// var genre3 int
		// row.Scan(&id, &user, &movie, &like, &watched, &watching, &toWatch, &genre1, &genre2, &genre3)
		row.Scan(&movie, &watching)
		// log.Println("Users: ", id, " ", user, " ", movie, " ", like, " ", watched, " ", watching, " ", toWatch, " ", genre1, " ", genre2, " ", genre3)
		userWatingSingle.userWatchingMovie = GetSingleMovie2(movie)
		userWatingSingle.userWatchingVal = watching
		userWatching = append(userWatching, userWatingSingle)
	}
	//************************************************************************************************************************************************************************

	//to watch
	row, err = db.Query("SELECT movieID FROM " + tablename2 + " WHERE idUser= " + strconv.Itoa(idUser) + " AND toWatch = 1")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		// var id int
		// var user int
		var movie int
		// var like int
		// var watched int
		// var watching int
		// var toWatch int
		// var genre1 int
		// var genre2 int
		// var genre3 int
		// row.Scan(&id, &user, &movie, &like, &watched, &watching, &toWatch, &genre1, &genre2, &genre3)
		row.Scan(&movie)
		// log.Println("Users: ", id, " ", user, " ", movie, " ", like, " ", watched, " ", watching, " ", toWatch, " ", genre1, " ", genre2, " ", genre3)

		userToWatch = append(userToWatch, GetSingleMovie2(movie))
	}
	//************************************************************************************************************************************************************************

	//watched
	row, err = db.Query("SELECT movieID FROM " + tablename2 + " WHERE idUser= " + strconv.Itoa(idUser) + " AND watched = 1")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		// var id int
		// var user int
		var movie int
		// var like int
		// var watched int
		// var watching int
		// var toWatch int
		// var genre1 int
		// var genre2 int
		// var genre3 int
		// row.Scan(&id, &user, &movie, &like, &watched, &watching, &toWatch, &genre1, &genre2, &genre3)
		row.Scan(&movie)
		// log.Println("Users: ", id, " ", user, " ", movie, " ", like, " ", watched, " ", watching, " ", toWatch, " ", genre1, " ", genre2, " ", genre3)

		userWatched = append(userWatched, GetSingleMovie2(movie))
	}

	//************************************************************************************************************************************************************************

	//Like
	row, err = db.Query("SELECT movieID, like FROM " + tablename2 + " WHERE idUser= " + strconv.Itoa(idUser) + " AND Like <> -1")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		// var id int
		// var user int
		var movie int
		var like int
		// var watched int
		// var watching int
		// var toWatch int
		// var genre1 int
		// var genre2 int
		// var genre3 int
		// row.Scan(&id, &user, &movie, &like, &watched, &watching, &toWatch, &genre1, &genre2, &genre3)
		row.Scan(&movie, &like)
		// log.Println("Users: ", id, " ", user, " ", movie, " ", like, " ", watched, " ", watching, " ", toWatch, " ", genre1, " ", genre2, " ", genre3)
		userLikeSingle.userLikeMovie = GetSingleMovie2(movie)
		userLikeSingle.userLikeVal = like
		userLike = append(userLike, userLikeSingle)
	}

	//************************************************************************************************************************************************************************
	//************************************************************************************************************************************************************************
	//************************************************************************************************************************************************************************

	c.JSON(http.StatusOK, gin.H{
		"Username": username,
		"Like":     userLike,
		"Watched":  userWatched,
		"Watching": userWatching,
		"ToWatch":  userToWatch,
	})

}

//************************************************************************************************************************************************************************

func GetSingleMovie2(movieID int) (filmId FilmId) {

	//Communication

	response, err := http.Get("https://api.themoviedb.org/3/movie/" + strconv.Itoa(movieID) + "?api_key=" + os.Getenv("Twilio_api_key") + "&language=en-US")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(responseData), &filmId)

	return

}