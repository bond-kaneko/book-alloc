package v1

import (
	"book-alloc/api/v1/handler"
	"book-alloc/internal/allocation"
	"book-alloc/internal/reading_history"
	"github.com/gin-gonic/gin"
)

func User(r *gin.RouterGroup) {
	r.GET("/users/me", handler.LoginUserHandler)
}

func Allocation(r *gin.RouterGroup) {
	r.GET("/allocations", allocation.GetAll)
}

func ReadingHistory(r *gin.RouterGroup) {
	r.GET("/reading-histories", reading_history.GetAll)
}
