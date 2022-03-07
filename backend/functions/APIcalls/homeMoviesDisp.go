package apiCall

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//************************************************************************************************************************************************************************

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

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		log.Fatal("2 Error loading .env file" + err.Error())
	}
}

//************************************************************************************************************************************************************************

func HomeMovies(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		strconv.Itoa(2022): GetMoviesOfAYear(2022),
		strconv.Itoa(2021): GetMoviesOfAYear(2021),
		strconv.Itoa(2020): GetMoviesOfAYear(2020),
		strconv.Itoa(2019): GetMoviesOfAYear(2019),
		strconv.Itoa(2018): GetMoviesOfAYear(2018),
		strconv.Itoa(2017): GetMoviesOfAYear(2017),
		strconv.Itoa(2016): GetMoviesOfAYear(2016),
		strconv.Itoa(2015): GetMoviesOfAYear(2015),
		strconv.Itoa(2014): GetMoviesOfAYear(2014),
		strconv.Itoa(2013): GetMoviesOfAYear(2013),
		strconv.Itoa(2012): GetMoviesOfAYear(2012),
	})

}

//************************************************************************************************************************************************************************
