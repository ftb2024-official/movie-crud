package middlewares

import (
	"io"

	"github.com/gin-gonic/gin"
)

func CheckReqJSON(ctx *gin.Context) {
	input, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Unable to read request body"})
		ctx.Abort()
	}
}
