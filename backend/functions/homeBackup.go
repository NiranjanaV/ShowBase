package functions

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

// type Film struct {
// 	Poster_path       string
// 	Adult             bool
// 	Overview          string
// 	Release_date      string
// 	Genre_ids         []int32
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

// type Api struct {
// 	Page          int
// 	Results       []Film
// 	Total_pages   int64
// 	Total_results int64
// }

// var api Api
// var film Film

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		log.Fatal("2 Error loading .env file" + err.Error())
	}
}

func HomeMovies2(c *gin.Context) {
	response, err := http.Get("https://api.themoviedb.org/3/discover/movie?api_key=" + os.Getenv("Twilio_api_key") + "&primary_release_year=2022&sort_by=revenue.desc")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(responseData), &api)
	fmt.Println(api)

	data, _ := json.Marshal(api.Results)
	fmt.Println(string(data))
	c.JSON(http.StatusOK, gin.H{
		"data": api.Results,
	})

}
