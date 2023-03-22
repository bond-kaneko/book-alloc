package main

import (
	"book-alloc/api/v1/allocation"
	"book-alloc/api/v1/user"
	"book-alloc/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	env := os.Getenv("ENV")
	if "" == env {
		env = "local"
	}
	if err := godotenv.Load(".env." + env); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("ORIGIN_WEB")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc:  nil,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "public"})
	})

	auth := r.Group("/auth", adapter.Wrap(middleware.EnsureValidToken()))
	{
		auth.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
		user.Handle(auth)
		allocation.Handle(auth)
	}

	r.Run()
}
