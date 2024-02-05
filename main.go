package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Director struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
}

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Year     int       `json:"year"`
	Director *Director `json:"director"`
}

var movies []Movie

func getMovies(db *[]Movie) gin.HandlerFunc {
	return func(ctx *gin.Context) { ctx.JSON(200, gin.H{"data": db}) }
}

func getMovie(db *[]Movie) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		found := false

		for _, movie := range *db {
			if movie.Id == id {
				ctx.JSON(200, gin.H{"data": movie})
				return
			}
		}

		if !found {
			ctx.JSON(404, gin.H{"Not Found": fmt.Sprintf("Movie with id = %v not found", id)})
		}
	}
}

func deleteMovie(db *[]Movie) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		idxToDelete := -1

		for idx, movie := range *db {
			if movie.Id == id {
				*db = append((*db)[:idx], (*db)[idx+1:]...)
				ctx.JSON(200, gin.H{"message": fmt.Sprintf("User with id = %v successfully deleted", id)})
				return
			}
		}

		if idxToDelete == -1 {
			ctx.JSON(404, gin.H{"Not Found": fmt.Sprintf("Movie with id = %v not found", id)})
		}
	}
}

func createMovie(db *[]Movie) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newMovie Movie

		input, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Unable to read request body"})
			return
		}

		err = json.Unmarshal(input, &newMovie)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Unable to marshall JSON request body"})
			return
		}

		newMovie.Id = uuid.NewString()
		*db = append(*db, newMovie)
		ctx.JSON(201, gin.H{"message": "New movie successfully created", "data": newMovie})
	}
}

func editMovie(db *[]Movie) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func main() {
	movies = append(movies, Movie{Id: uuid.NewString(), Isbn: "227125917-7", Title: "Titanic", Year: 1999, Director: &Director{FirstName: "John", LastName: "Doe"}})
	router := gin.Default()

	router.GET("/api/movies", getMovies(&movies))
	router.GET("/api/movies/:id", getMovie(&movies))
	router.POST("/api/movies", createMovie(&movies))
	router.PATCH("/api/movies/:id", editMovie(&movies))
	router.DELETE("/api/movies/:id", deleteMovie(&movies))

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Unable to run the server at port 8080 :(")
	}
}
