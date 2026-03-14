package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/s2a-go/retry"
)

func AuthMiddleware() gin.HandlerFunc {
	return func (c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "gak nemu Authorization",
				"error_code": "MISSING_TOKEN",
			})
			return
		}
	}
}

