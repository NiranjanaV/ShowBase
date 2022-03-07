package apiCall

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//************************************************************************************************************************************************************************

//chnage movie struct if needed ,,,its different for queries
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

type Genre struct {
	Id   int32
	Name string
}

var filmId FilmId

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		log.Fatal("7 Error loading .env file" + err.Error())
	}
}

//************************************************************************************************************************************************************************

func GetMovie(c *gin.Context) {

	//Communication
	movieID := c.Param("movie")

	response, err := http.Get("https://api.themoviedb.org/3/movie/" + movieID + "?api_key=" + os.Getenv("Twilio_api_key") + "&language=en-US")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(responseData), &filmId)

	if len(filmId.Genres) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"movie": filmId,
		})
	} else if len(filmId.Genres) == 1 {
		c.JSON(http.StatusOK, gin.H{
			"movie":               filmId,
			filmId.Genres[0].Name: GetMoviesOfAGenra(filmId.Genres[0]),
		})
	} else if len(filmId.Genres) == 2 {
		c.JSON(http.StatusOK, gin.H{
			"movie":               filmId,
			filmId.Genres[0].Name: GetMoviesOfAGenra(filmId.Genres[0]),
			filmId.Genres[1].Name: GetMoviesOfAGenra(filmId.Genres[1]),
		})
	} else if len(filmId.Genres) == 3 {
		c.JSON(http.StatusOK, gin.H{
			"movie":               filmId,
			filmId.Genres[0].Name: GetMoviesOfAGenra(filmId.Genres[0]),
			filmId.Genres[1].Name: GetMoviesOfAGenra(filmId.Genres[1]),
			filmId.Genres[2].Name: GetMoviesOfAGenra(filmId.Genres[2]),
		})
	} else if len(filmId.Genres) == 4 {
		c.JSON(http.StatusOK, gin.H{
			"movie":               filmId,
			filmId.Genres[0].Name: GetMoviesOfAGenra(filmId.Genres[0]),
			filmId.Genres[1].Name: GetMoviesOfAGenra(filmId.Genres[1]),
			filmId.Genres[2].Name: GetMoviesOfAGenra(filmId.Genres[2]),
			filmId.Genres[3].Name: GetMoviesOfAGenra(filmId.Genres[3]),
		})
	} else if len(filmId.Genres) == 5 {
		c.JSON(http.StatusOK, gin.H{
			"movie":               filmId,
			filmId.Genres[0].Name: GetMoviesOfAGenra(filmId.Genres[0]),
			filmId.Genres[1].Name: GetMoviesOfAGenra(filmId.Genres[1]),
			filmId.Genres[2].Name: GetMoviesOfAGenra(filmId.Genres[2]),
			filmId.Genres[3].Name: GetMoviesOfAGenra(filmId.Genres[3]),
			filmId.Genres[4].Name: GetMoviesOfAGenra(filmId.Genres[4]),
		})
	} else if len(filmId.Genres) == 6 {
		c.JSON(http.StatusOK, gin.H{
			"movie":               filmId,
			filmId.Genres[0].Name: GetMoviesOfAGenra(filmId.Genres[0]),
			filmId.Genres[1].Name: GetMoviesOfAGenra(filmId.Genres[1]),
			filmId.Genres[2].Name: GetMoviesOfAGenra(filmId.Genres[2]),
			filmId.Genres[3].Name: GetMoviesOfAGenra(filmId.Genres[3]),
			filmId.Genres[4].Name: GetMoviesOfAGenra(filmId.Genres[4]),
			filmId.Genres[5].Name: GetMoviesOfAGenra(filmId.Genres[5]),
		})
	} else if len(filmId.Genres) == 7 {
		c.JSON(http.StatusOK, gin.H{
			"movie":               filmId,
			filmId.Genres[0].Name: GetMoviesOfAGenra(filmId.Genres[0]),
			filmId.Genres[1].Name: GetMoviesOfAGenra(filmId.Genres[1]),
			filmId.Genres[2].Name: GetMoviesOfAGenra(filmId.Genres[2]),
			filmId.Genres[3].Name: GetMoviesOfAGenra(filmId.Genres[3]),
			filmId.Genres[4].Name: GetMoviesOfAGenra(filmId.Genres[4]),
			filmId.Genres[5].Name: GetMoviesOfAGenra(filmId.Genres[5]),
			filmId.Genres[6].Name: GetMoviesOfAGenra(filmId.Genres[6]),
		})
	} else if len(filmId.Genres) == 8 {
		c.JSON(http.StatusOK, gin.H{
			"movie":               filmId,
			filmId.Genres[0].Name: GetMoviesOfAGenra(filmId.Genres[0]),
			filmId.Genres[1].Name: GetMoviesOfAGenra(filmId.Genres[1]),
			filmId.Genres[2].Name: GetMoviesOfAGenra(filmId.Genres[2]),
			filmId.Genres[3].Name: GetMoviesOfAGenra(filmId.Genres[3]),
			filmId.Genres[4].Name: GetMoviesOfAGenra(filmId.Genres[4]),
			filmId.Genres[5].Name: GetMoviesOfAGenra(filmId.Genres[5]),
			filmId.Genres[6].Name: GetMoviesOfAGenra(filmId.Genres[6]),
			filmId.Genres[7].Name: GetMoviesOfAGenra(filmId.Genres[7]),
		})
	} else if len(filmId.Genres) >= 9 {
		c.JSON(http.StatusOK, gin.H{
			"movie":               filmId,
			filmId.Genres[0].Name: GetMoviesOfAGenra(filmId.Genres[0]),
			filmId.Genres[1].Name: GetMoviesOfAGenra(filmId.Genres[1]),
			filmId.Genres[2].Name: GetMoviesOfAGenra(filmId.Genres[2]),
			filmId.Genres[3].Name: GetMoviesOfAGenra(filmId.Genres[3]),
			filmId.Genres[4].Name: GetMoviesOfAGenra(filmId.Genres[4]),
			filmId.Genres[5].Name: GetMoviesOfAGenra(filmId.Genres[5]),
			filmId.Genres[6].Name: GetMoviesOfAGenra(filmId.Genres[6]),
			filmId.Genres[7].Name: GetMoviesOfAGenra(filmId.Genres[7]),
			filmId.Genres[8].Name: GetMoviesOfAGenra(filmId.Genres[8]),
		})
	}

}

//************************************************************************************************************************************************************************
