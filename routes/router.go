package routes

import (
	"github.com/RIOKOWI/tugas_week_5_1123150059/handlers"
	"github.com/RIOKOWI/tugas_week_5_1123150059/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
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
	cartHandler := handlers.NewCartHandler()
	orderHandler := handlers.NewOrderHandler()


	v1 := r.Group("/v1")
	{
	
			v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok", "service": "rio-achyar"})
		})

	
	auth := v1.Group("/auth")
	{
		auth.POST("/verify-token", authHandler.VerifyToken)
	}

	protected := v1.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		products := protected.Group("/products")
	{
			products.GET("", productHandler.GetAll) 
			products.GET("/:id", productHandler.GetByID) 
		
			adminProducts := products.Group("")
			adminProducts.Use(middleware.AdminOnly())
				{
					adminProducts.POST("", productHandler.Create) 
					adminProducts.PUT("/:id", productHandler.Update) 
					adminProducts.DELETE("/:id", productHandler.Delete)
				}
			}
			// Cart
			cart := protected.Group("/cart")
			{
				cart.GET("", cartHandler.GetCart)           // GET    /v1/cart
				cart.POST("", cartHandler.AddToCart)        // POST   /v1/cart
				cart.PUT("/:id", cartHandler.UpdateCartItem) // PUT    /v1/cart/:id
				cart.DELETE("/:id", cartHandler.RemoveCartItem) // DELETE /v1/cart/:id
				cart.DELETE("", cartHandler.ClearCart)      // DELETE /v1/cart
			}

			// Orders
			orders := protected.Group("/orders")
			{
				orders.POST("/checkout", orderHandler.Checkout)    // POST   /v1/orders/checkout
				orders.GET("", orderHandler.GetMyOrders)           // GET    /v1/orders
				orders.GET("/:id", orderHandler.GetOrderByID)      // GET    /v1/orders/:id
			}

			// Admin — order management
			admin := protected.Group("/admin")
			admin.Use(middleware.AdminOnly())
			{
				admin.GET("/orders", orderHandler.GetAllOrders)                     // GET /v1/admin/orders
				admin.PUT("/orders/:id/status", orderHandler.UpdateOrderStatus)     // PUT /v1/admin/orders/:id/status
			}
		}
	}
	return r
}
