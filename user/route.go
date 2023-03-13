package user

import (
	"github.com/gin-gonic/gin"
)

func Route(r *gin.RouterGroup) {
	r.GET("/users", GetAll)
}
