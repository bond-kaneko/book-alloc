package main

import (
	"book-alloc/api/v1"
	"book-alloc/internal/middleware"
	"book-alloc/internal/user"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
)

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	u, _ := c.Get(middleware.IdentityKey)
	c.JSON(200, gin.H{
		"userID":   claims[middleware.IdentityKey],
		"userName": u.(*user.User).Email,
		"text":     "Hello World.",
	})
}

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
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)

		v1.User(auth)
		v1.Allocation(auth)
		v1.ReadingHistory(auth)
	}

	g.Run()
}
