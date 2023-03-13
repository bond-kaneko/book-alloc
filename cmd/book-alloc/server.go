package main

import (
	"book-alloc/user"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	routes := g.Group("/v1")
	{
		user.Route(routes)
	}

	g.Run()
}
