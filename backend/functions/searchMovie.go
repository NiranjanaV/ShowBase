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
		log.Fatal("4 Error loading .env files" + err.Error())
	}
}

func Search(text string) {
	response, err := http.Get("https://api.themoviedb.org/3/search/movie?api_key=" + os.Getenv("Twilio_api_key") + "&language=en-US&query=" + text + "&page=1")

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

	pages, _ := json.Marshal(api.Total_pages)
	fmt.Println(string(pages))

	data, _ := json.Marshal(api.Results)
	fmt.Println(string(data))
}

func SearchWithPage(text string, page int) {
	response, err := http.Get("https://api.themoviedb.org/3/search/movie?api_key=" + os.Getenv("Twilio_api_key") + "&language=en-US&query=" + text + "&page=" + strconv.Itoa(page))

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

	pages, _ := json.Marshal(api.Total_pages)
	fmt.Println(string(pages))

	data, _ := json.Marshal(api.Results)
	fmt.Println(string(data))
}
