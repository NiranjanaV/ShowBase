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

//************************************************************************************************************************************************************************

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		log.Fatal("7 Error loading .env file" + err.Error())
	}
}

//************************************************************************************************************************************************************************

func GetMoviesOfAGenra(genre Genre) (api Api) {
	//Communication
	genId := genre.Id
	response, err := http.Get("https://api.themoviedb.org/3/discover/movie?api_key=" + os.Getenv("Twilio_api_key") + "&language=en-US&sort_by=popularity.desc&page=1&with_genres=" + strconv.Itoa(int(genId)) + "&with_watch_monetization_types=flatrate")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(responseData), &api)

	return
}

//************************************************************************************************************************************************************************

func GetMoviesOfAGenraId(genId int32) (api Api) {
	//Communication
	response, err := http.Get("https://api.themoviedb.org/3/discover/movie?api_key=" + os.Getenv("Twilio_api_key") + "&language=en-US&sort_by=popularity.desc&page=1&with_genres=" + strconv.Itoa(int(genId)) + "&with_watch_monetization_types=flatrate")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(responseData), &api)

	return
}
