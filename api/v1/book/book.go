package book

import (
	"github.com/gin-gonic/gin"
)

func Handle(r *gin.RouterGroup) {
	a := r.Group("/books")
	{
		a.GET("/:userId", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ok",
			})
		})
	}
}
