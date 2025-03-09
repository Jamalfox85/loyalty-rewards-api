package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://api.bigcommerce.com"}, // Allow only your frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Test Route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Service is running!"})
	})
	r.GET("/auth", func(c *gin.Context) {
		// c.JSON(http.StatusOK, gin.H{"message": "CORS is working!"})
		c.Redirect(http.StatusFound, "https://supposedly-simple-hamster.ngrok-free.app")
	})
	r.GET("/load", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "CORS is working!"})
	})
	r.GET("/uninstall", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "CORS is working!"})
	})
	r.GET("/remove-user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "CORS is working!"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Default port if not set
	}
	r.Run("0.0.0.0:" + port) // Start server
}