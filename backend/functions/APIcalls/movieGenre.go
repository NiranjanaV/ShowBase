package apiCall

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

//************************************************************************************************************************************************************************

var genr []Genre

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		//fmt.Println("6 Error loading .env file" + err.Error())
	}
}

//************************************************************************************************************************************************************************

func GetGenre(movieID int) (gen []int) {

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
