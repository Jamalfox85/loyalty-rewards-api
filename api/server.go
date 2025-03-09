package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamalfox85/loyalty-rewards-api/api/handlers"
)

type Server struct {
	listenAddr string
}

func NewServer(	listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start(app *Application) {
	router := gin.Default()
	router.Use(CORSMiddleware())

	// Users
	router.GET("/plans", handlers.FetchPlan(app.Plans))
	
	// Test Route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Service is running!"})
	})

	// BC Redirect Routes
	router.GET("/auth", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://supposedly-simple-hamster.ngrok-free.app")
	})
	router.GET("/load", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://supposedly-simple-hamster.ngrok-free.app")
	})
	router.GET("/uninstall", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://supposedly-simple-hamster.ngrok-free.app")
	})
	router.GET("/remove-user", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://supposedly-simple-hamster.ngrok-free.app")
	})


	fmt.Println("Server Running on", s.listenAddr);
	if err := router.Run(":" + s.listenAddr); err != nil {
		log.Panicf("error: %s", err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}