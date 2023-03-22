package allocation

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateRequest struct {
	UserId   string
	Name     string
	Share    int
	IsActive bool
}

func Handle(r *gin.RouterGroup) {
	a := r.Group("/allocations")
	{
		a.POST("/", handleCreate)
	}
}

func handleCreate(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
