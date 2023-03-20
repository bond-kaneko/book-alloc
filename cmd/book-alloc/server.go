package main

import (
	"book-alloc/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc:  nil,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "public"})
	})

	auth := r.Group("/auth", adapter.Wrap(middleware.EnsureValidToken()))
	auth.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.Run()
}
