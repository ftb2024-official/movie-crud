package middlewares

import (
	"encoding/json"
	"io"

	movie "github.com/ftb2024-official/movie-crud/entity"
	"github.com/gin-gonic/gin"
)

func CheckReqJSON(ctx *gin.Context) {
	input, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Unable to read request body"})
		return
	}

	var newMovie movie.Movie

	if err2 := json.Unmarshal(input, &newMovie); err2 != nil {
		ctx.JSON(400, gin.H{"error": "Unable to marshall JSON request body"})
		return
	}

	if newMovie.Title == "" || !(len(newMovie.Title) > 2) {
		ctx.JSON(400, gin.H{"error": "Title is required | At least 3 characters for Title"})
		return
	}

	if newMovie.Year <= 1900 || newMovie.Year >= 2024 {
		ctx.JSON(400, gin.H{"error": "Year is required | Year must be between 1900 and 2024"})
		return
	}

	if newMovie.Director.FirstName == "" || !(len(newMovie.Director.FirstName) > 2) {
		ctx.JSON(400, gin.H{"error": "First Name is required | At least 3 characters for First Name"})
		return
	}

	if newMovie.Director.LastName == "" || !(len(newMovie.Director.LastName) > 2) {
		ctx.JSON(400, gin.H{"error": "Last Name is required | At least 3 characters for Last Name"})
		return
	}

	ctx.Set("validJSON", newMovie)
	ctx.Next()
}
