package v1

import (
	"book-alloc/internal/allocation"
	"book-alloc/internal/reading_history"
	"book-alloc/internal/user"
	"github.com/gin-gonic/gin"
)

func User(r *gin.RouterGroup) {
	r.GET("/users", user.GetAll)
}

func Allocation(r *gin.RouterGroup) {
	r.GET("/allocations", allocation.GetAll)
}

func ReadingHistory(r *gin.RouterGroup) {
	r.GET("/reading-histories", reading_history.GetAll)
}
