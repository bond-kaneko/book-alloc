package main

import (
	"book-alloc/api/v1"
	"book-alloc/internal"
	"book-alloc/internal/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	g.Use(sessions.Sessions("mysession", store))
	g.POST("/login", internal.Login)

	routes := g.Group("/v1")
	routes.Use(middleware.LoginCheckMiddleware())
	{
		v1.User(routes)
		v1.Allocation(routes)
		v1.ReadingHistory(routes)
	}

	g.Run()
}
