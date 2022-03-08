package main

import (
	D "main/dbSQLite"
	A "main/functions/APIcalls"
	DM "main/functions/dbAccess"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//************************************************************************************************************************************************************************

func main() {

	// Communication from frontend
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Functions to be called from the client
	r.GET("/search/:name", A.Search)
	r.GET("/top", A.HomeMovies)
	r.GET("/searchPage", A.SearchWithPage)
	r.PUT("/userReg", DM.InsertAuthTable)
	r.GET("/displayPass", DM.DisplayAuthTable)
	r.PUT("/authenticateUser", DM.GetPassForUser)
	r.GET("/getMovie/:movie", A.GetMovie)
	r.GET("/userLikes/:username", DM.GetUserTable)
	r.PUT("/updateUserLike", DM.InsertUserTable)
	r.GET("/displayLike", DM.DisplayUserTable)
	r.GET("/getGenre/:GenreId", A.SearchGenre)
	r.GET("/getGenraPage", A.SearchGenreWithPage)
	r.GET("/getUserProfile/:username", DM.SendUserProfile)

	//initializeRoutes()
	err := r.Run()
	if err != nil {
		panic("Failed to run the server")
	}
	db := D.GetDB()
	defer db.Close()
}

//************************************************************************************************************************************************************************
