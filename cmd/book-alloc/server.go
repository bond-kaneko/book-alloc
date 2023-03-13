package main

import (
	"book-alloc/api/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	routes := g.Group("/v1")
	{
		v1.User(routes)
	}

	g.Run()
}
