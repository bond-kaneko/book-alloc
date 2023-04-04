package router

import (
	"book-alloc/api/handler/allocation_handler"
	"book-alloc/api/handler/reading_experience_handler"
	"book-alloc/api/handler/user_handler"
	"book-alloc/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	"net/http"
	"os"
	"time"
)

func Initialize() *gin.Engine {
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
		user_handler.Routes(auth)
		allocation_handler.Routes(auth)
		reading_experience_handler.Routes(auth)
	}

	return r
}
