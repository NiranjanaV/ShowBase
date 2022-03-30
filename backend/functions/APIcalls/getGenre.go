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

type GenreList struct {
	Genres []Genre
}

var genreList GenreList

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		log.Fatal("7 Error loading .env file" + err.Error())
	}
}

//************************************************************************************************************************************************************************

func GetAllGenre(c *gin.Context) {

	//Communication
	response, err := http.Get("https://api.themoviedb.org/3/genre/movie/list?api_key=" + os.Getenv("Twilio_api_key") + "&language=en-US")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(responseData), &genreList)

	c.JSON(http.StatusOK, gin.H{
		"Genres": genreList,
	})
}

//************************************************************************************************************************************************************************
