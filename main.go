package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	movie "github.com/ftb2024-official/movie-crud/entity"
	middleware "github.com/ftb2024-official/movie-crud/middlewares"
)

var movies []movie.Movie

func getMovies(db *[]movie.Movie) gin.HandlerFunc {
	return func(ctx *gin.Context) { ctx.JSON(200, gin.H{"data": db}) }
}

func getMovie(db *[]movie.Movie) gin.HandlerFunc {
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

func deleteMovie(db *[]movie.Movie) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		idxToDelete := -1

		for idx, movie := range *db {
			if movie.Id == id {
				*db = append((*db)[:idx], (*db)[idx+1:]...)
				ctx.JSON(200, gin.H{"message": fmt.Sprintf("Movie with id = %v successfully deleted", id)})
				return
			}
		}

		if idxToDelete == -1 {
			ctx.JSON(404, gin.H{"Not Found": fmt.Sprintf("Movie with id = %v not found", id)})
		}
	}
}

func createMovie(db *[]movie.Movie) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newMovie movie.Movie

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

func editMovie(db *[]movie.Movie) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var editedMovie, oldMovie movie.Movie
		found := false

		id := ctx.Param("id")

		for _, movie := range *db {
			if movie.Id == id {
				found = true
				oldMovie = movie
				break
			}
		}

		if !found {
			ctx.JSON(404, gin.H{"Not Found": fmt.Sprintf("Movie with id = %v not found", id)})
			return
		}

		input, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Unable to read request body"})
			return
		}

		err = json.Unmarshal(input, &editedMovie)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Unable to marshall JSON request body"})
			return
		}

		if editedMovie.Isbn != "" {
			oldMovie.Isbn = editedMovie.Isbn
		}

		if editedMovie.Title != "" {
			oldMovie.Title = editedMovie.Title
		}

		if editedMovie.Year != 0 {
			oldMovie.Year = editedMovie.Year
		}

		if editedMovie.Isbn != "" {
			oldMovie.Isbn = editedMovie.Isbn
		}

		if editedMovie.Director.FirstName != "" {
			oldMovie.Director.FirstName = editedMovie.Director.FirstName
		}

		if editedMovie.Director.LastName != "" {
			oldMovie.Director.LastName = editedMovie.Director.LastName
		}

		ctx.JSON(201, gin.H{
			"message": fmt.Sprintf("Movie with id = %v successfully edited", id),
			"data":    oldMovie,
		})
	}
}

func main() {
	movies = append(movies, movie.Movie{Id: uuid.NewString(), Isbn: uuid.NewString(), Title: "Titanic", Year: 1999, Director: movie.Director{FirstName: "John", LastName: "Doe"}})
	router := gin.Default()

	router.GET("/api/movies", getMovies(&movies))
	router.Use(middleware.CheckReqJSON).POST("/api/movies", createMovie(&movies))
	router.Use(middleware.CheckGetUUID).GET("/api/movies/:id", getMovie(&movies))
	router.Use(middleware.CheckEditUUID).PATCH("/api/movies/:id", editMovie(&movies))
	router.Use(middleware.ChecDelUUID).DELETE("/api/movies/:id", deleteMovie(&movies))

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Unable to run the server at port 8080 :(")
	}
}
