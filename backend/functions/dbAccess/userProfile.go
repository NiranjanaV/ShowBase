package tables

import (
	"fmt"
	"log"

	// D "main/dbSQLite"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func init() {

	err := godotenv.Load("go.env")

	if err != nil {
		log.Fatal("1 Error loading .env file" + err.Error())
	}
}

//************************************************************************************************************************************************************************

func sendUserProfile(c *gin.Context) {

	idUser := GetUserID(c.Param("username"))
	fmt.Println(idUser)
	fmt.Println("disp")
	row, err := db.Query("SELECT * FROM " + tablename2 + " WHERE idUser= " + strconv.Itoa(idUser))
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var user int
		var movie int
		var like int
		var watched int
		var watching int
		var toWatch int
		var genre1 int
		var genre2 int
		var genre3 int
		row.Scan(&id, &user, &movie, &like, &watched, &watching, &toWatch, &genre1, &genre2, &genre3)
		log.Println("Users: ", id, " ", user, " ", movie, " ", like, " ", watched, " ", watching, " ", toWatch, " ", genre1, " ", genre2, " ", genre3)
		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"movie":    movie,
			"user":     user,
			"like":     like,
			"watched":  watched,
			"watching": watching,
			"toWatch":  toWatch,
			"genre1":   genre1,
			"genre2":   genre2,
			"genre3":   genre3,
		})
	}
}

//************************************************************************************************************************************************************************
