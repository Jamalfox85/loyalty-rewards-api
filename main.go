package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Allow only your frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Test Route
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "CORS is working!"})
	})
	r.GET("/oauth", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "CORS is working!"})
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

	r.Run(":8080") // Start server
}