package v1

import (
	"book-alloc/internal/allocation"
	"book-alloc/internal/reading_history"
	"github.com/gin-gonic/gin"
	"net/http"
)

func User(r *gin.RouterGroup) {
	u := r.Group("/users")
	{
		u.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
}

func Allocation(r *gin.RouterGroup) {
	r.GET("/allocations", allocation.GetAll)
}

func ReadingHistory(r *gin.RouterGroup) {
	r.GET("/reading-histories", func(c *gin.Context) {
		rh := reading_history.GetAll()
		c.JSON(http.StatusOK, rh)
	})
}
