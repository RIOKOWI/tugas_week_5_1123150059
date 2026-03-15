package routes

import (
	"github.com/RIOKOWI/tugas_week_5_1123150059/handlers"
	"github.com/RIOKOWI/tugas_week_5_1123150059/middleware"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	authHandler := handlers.NewAuthHandler()
	productHandler := handlers.NewProductHandler()

	v1 := r.Group("/v1")
	{
	
			v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok", "service": "rio-achyar"})
		})

	}
	auth := v1.Group("/auth")
	{
		auth.POST("/verify-token", authHandler.VerifyToken)
	}
}
