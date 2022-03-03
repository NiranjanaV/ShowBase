package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Film struct {
	Poster_path       string
	Adult             bool
	Overview          string
	Release_date      string
	Genre_ids         []int32
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

type Api struct {
	Page          int
	Results       []Film
	Total_pages   int64
	Total_results int64
}

type Json struct {
	Title   string
	Results []Film
}

var api Api
var film Film
var json2 Json

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		log.Fatal("2 Error loading .env file" + err.Error())
	}
}

func HomeMovies() {

	movies := make([]Film, 10)
	titles := make([]string, 10)
	for i := 0; i < 10; i++ {

		response, err := http.Get("https://api.themoviedb.org/3/discover/movie?api_key=" + os.Getenv("Twilio_api_key") + "&primary_release_year=" + strconv.Itoa(2022-i) + "&sort_by=revenue.desc")

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

		json.Unmarshal([]byte(data), &film)
		fmt.Println(data)

		movies[i] = film
		titles[i] = strconv.Itoa(2022 - i)
	}

	jsons, err := json.Marshal(movies)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(jsons))
	}
	// data, _ := json.Marshal(api.Results)
	// fmt.Println(string(data))

}
