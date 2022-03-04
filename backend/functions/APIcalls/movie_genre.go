package apiCall

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

// //chnage movie struct if needed ,,,its different for queries
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
var genr []Genre

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		log.Fatal("3 Error loading .env file" + err.Error())
	}
}

func GetGenre(movieID int) (gen []int) {

	// TO be received by user
	// movieID := 332562

	//Code
	response, err := http.Get("https://api.themoviedb.org/3/movie/" + strconv.Itoa(movieID) + "?api_key=" + os.Getenv("Twilio_api_key") + "&language=en-US")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// fmt.Println(response)

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(responseData)

	json.Unmarshal([]byte(responseData), &filmId)
	// fmt.Println(filmId)

	// json.Unmarshal([]byte(filmId.Genres), &genr)
	// fmt.Println(genr)

	data, _ := json.Marshal(filmId.Genres)

	json.Unmarshal([]byte(data), &genr)
	// fmt.Println(genr)

	// for genres := range genr {
	// 	ids, _ := json.Marshal(genres.Id)
	// 	fmt.Println(string(ids))
	// }
	gen = []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	for i, s := range genr {
		var g int
		genres, _ := json.Marshal(s.Id)
		json.Unmarshal([]byte(genres), &g)
		gen[i] = g
	}
	fmt.Println(gen)
	// genres, _ := json.Marshal(genr)
	// fmt.Println(string(genres))

	// fmt.Println(string(data))
	// handleRequestsGenre()
	return
}
