package book

import (
	"book-alloc/db"
	"book-alloc/internal/book"
	"github.com/gin-gonic/gin"
)

func Handle(r *gin.RouterGroup) {
	a := r.Group("/books")
	{
		a.GET("/:userId", HandleMyBooks)
	}
}

func HandleMyBooks(c *gin.Context) {
	userId := c.Param("userId")
	d, _ := db.NewDB()

	books := book.GetMyBooks(d, userId)

	c.JSON(200, books)
}
