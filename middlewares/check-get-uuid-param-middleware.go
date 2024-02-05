package middlewares

import (
	"github.com/gin-gonic/gin"

	util "github.com/ftb2024-official/movie-crud/utils"
)

func CheckGetUUID(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(400, gin.H{"error": "Parameter `id` is required"})
		ctx.Abort()
	}

	if util.IsUUID(id) {
		ctx.Next()
	} else {
		ctx.JSON(400, gin.H{"error": "Parameter `id` is not uuid"})
		ctx.Abort()
	}
}
