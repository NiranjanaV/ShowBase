package apiCall

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//************************************************************************************************************************************************************************

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		//fmt.Println("4 Error loading .env files" + err.Error())
	}
}

//************************************************************************************************************************************************************************

func SearchGenre(c *gin.Context) {
	Genreid := c.Param("GenreId")

	// //fmt.Println(reflect.ValueOf(Genreid), " ", Genreid)
	response, err := http.Get("https://api.themoviedb.org/3/discover/movie?api_key=" + os.Getenv("Twilio_api_key") + "&language=en-US&sort_by=popularity.desc&page=1&with_genres=" + Genreid + "&with_watch_monetization_types=flatrate")

	if err != nil {
		//fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//fmt.Println(err)
	}

	json.Unmarshal([]byte(responseData), &api)
	//fmt.Println(api)

	// pages, _ := json.Marshal(api.Total_pages)
	//fmt.Println(string(pages))

	c.JSON(http.StatusOK, gin.H{
		"Movies": api,
	})
}

//************************************************************************************************************************************************************************

func SearchGenreWithPage(c *gin.Context) {

	type GenrePage struct {
		Genreid string `json:"Genreid"`
		Page    string `json:"page"`
	}
	genrePage := GenrePage{}
	err := c.ShouldBindJSON(&genrePage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters, password should be between 8 to 20 chars",
		})
		return
	}

	response, err := http.Get("https://api.themoviedb.org/3/discover/movie?api_key=" + os.Getenv("Twilio_api_key") + "&language=en-US&sort_by=popularity.desc&page=" + genrePage.Page + "&with_genres=" + genrePage.Genreid + "&with_watch_monetization_types=flatrate")

	if err != nil {
		//fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//fmt.Println(err)
	}

	json.Unmarshal([]byte(responseData), &api)
	//fmt.Println(api)

	// pages, _ := json.Marshal(api.Total_pages)
	//fmt.Println(string(pages))

	// data, _ := json.Marshal(api.Results)
	//fmt.Println(string(data))

	c.JSON(http.StatusOK, gin.H{
		"Movies": api,
	})
}

//************************************************************************************************************************************************************************
