package middleware

import (
	"net/http"
	"strings"

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

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"succcess": false,
				"message": "format tokeb salah. gunakan: Bearer <token>",
				"error_code": "INVALID_TOKEN_FORMAT",
			})
			return
		}
	}
}

