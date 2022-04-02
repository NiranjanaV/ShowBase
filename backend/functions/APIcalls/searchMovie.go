package apiCall

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//************************************************************************************************************************************************************************

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		log.Fatal("4 Error loading .env files" + err.Error())
	}
}

//************************************************************************************************************************************************************************

func Search(c *gin.Context) {
	text := c.Param("name")

	fmt.Println(reflect.ValueOf(text), " ", text)
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
	// fmt.Println(api)

	// pages, _ := json.Marshal(api.Total_pages)
	// fmt.Println(string(pages))

	c.JSON(http.StatusOK, gin.H{
		"Movies": api,
	})
}

//************************************************************************************************************************************************************************

func SearchWithPage(c *gin.Context) {

	type SearchPage struct {
		Text string `json:"text"`
		Page string `json:"page"`
	}
	searchPage := SearchPage{}
	err := c.ShouldBindJSON(&searchPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters, password should be between 8 to 20 chars",
		})
		return
	}

	response, err := http.Get("https://api.themoviedb.org/3/search/movie?api_key=" + os.Getenv("Twilio_api_key") + "&language=en-US&query=" + searchPage.Text + "&page=" + searchPage.Page)

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

	c.JSON(http.StatusOK, gin.H{
		"Movies": api,
	})
}

//************************************************************************************************************************************************************************
