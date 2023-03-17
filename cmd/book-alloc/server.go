package main

import (
	"book-alloc/api/v1"
	"book-alloc/internal/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	g := gin.Default()

	authMiddleware, err := middleware.NewJwtMiddleware()
	if err != nil {
		logrus.Error("setUp auth failed: ?", err)
	}
	g.POST("/login", authMiddleware.LoginHandler)
	g.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := g.Group("/auth")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		route := auth.Group("/v1")
		v1.User(route)
		v1.Allocation(route)
		v1.ReadingHistory(route)
	}

	g.Run()
}
