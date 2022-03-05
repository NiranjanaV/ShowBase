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
