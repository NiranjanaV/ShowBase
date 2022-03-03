package main

import (
	O "main/functions"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var r *gin.Engine

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// r.GET("/searchPlaces/:name", showSearchPlacesPage)
	r.GET("/home", O.HomeMovies2)
	// r.POST("/login", userLogin)
	// r.GET("/users", getallUsers)
	// r.GET("/userprofile", getallTouristprofile)
	// r.GET("/guideprofile", getallGuideprofile)
	// r.GET("/comments", getallComments)
	// r.GET("/users/:email", getUser)
	// r.GET("/userprofile/:email", getTouristProfile)
	// r.GET("/guideprofile/:email", getGuideProfile)
	// r.GET("/comments/:location", getLocationComments)
	// r.POST("/userprofile", createTouristProfile)
	// r.POST("/guideprofile", createGuideProfile)
	// r.POST("/comments", createComments)
	// r.PUT("/userprofile/:email", updateTouristProfile)
	// r.PUT("/guideprofile/:email", updateGuideProfile)
	// r.DELETE("/userprofile/:email", DeleteTouristProfile)
	// r.DELETE("/guideprofile/:email", DeleteGuideProfile)

	//initializeRoutes()
	err := r.Run()
	if err != nil {
		panic("Failed to run the server")
	}

}
