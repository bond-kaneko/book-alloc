package allocation

import (
	"book-alloc/db"
	"book-alloc/internal/allocation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateRequest struct {
	UserId   string
	Name     string
	Share    int
	IsActive bool
}

func (cr *CreateRequest) toAllocation() allocation.Allocation {
	return allocation.Allocation{
		UserId:   cr.UserId,
		Name:     cr.Name,
		Share:    cr.Share,
		IsActive: cr.IsActive,
	}
}

func Handle(r *gin.RouterGroup) {
	a := r.Group("/allocations")
	{
		a.POST("/", handleCreate)
	}
}

func handleCreate(c *gin.Context) {
	var request CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	d, err := db.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dt := d.Begin()
	err = allocation.Create(dt, request.toAllocation())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	u, exists := allocation.GetLatestByUserId(dt, request.UserId)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dt.Commit()

	c.JSON(http.StatusOK, u)
}
