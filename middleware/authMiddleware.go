package middleware

import (
	"net/http"

	"github.com/DevanshBatra20/Password-Manager/helpers"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := helpers.ExtractToken(ctx)
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized Request",
				"data":    "Invalid request",
			})
			return
		}

		_, err := helpers.ValidateToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "Unauthorized",
				"data":   err.Error(),
			})
			return
		}

		ctx.Next()
	}
}
