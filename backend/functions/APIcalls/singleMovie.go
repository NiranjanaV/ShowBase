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

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		//fmt.Println("7 Error loading .env file" + err.Error())
	}
}

//************************************************************************************************************************************************************************

func GetSingleMovie(movieID int) (filmId FilmId) {

	//Communication

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

	return

}

//************************************************************************************************************************************************************************
