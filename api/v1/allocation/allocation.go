package allocation

import (
	"book-alloc/db"
	"book-alloc/internal/allocation"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Handle(r *gin.RouterGroup) {
	a := r.Group("/allocations")
	{
		a.POST("/", handleCreate)
		a.PUT("/", handleUpdate)
		a.GET("/:userId", handleGetByUserId)
	}
}

type createRequest struct {
	UserId   string `json:"UserId" binding:"required"`
	Name     string `json:"Name" binding:"required"`
	Share    int    `json:"Share" binding:"required"`
	IsActive bool   `json:"IsActive" binding:"required"`
}

func (cr *createRequest) toAllocation() allocation.Allocation {
	return allocation.Allocation{
		UserId:   cr.UserId,
		Name:     cr.Name,
		Share:    cr.Share,
		IsActive: cr.IsActive,
	}
}

func handleCreate(c *gin.Context) {
	var request createRequest
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

func handleGetByUserId(c *gin.Context) {
	d, err := db.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userId := c.Param("userId")
	u := allocation.GetByUserId(d, userId)

	c.JSON(http.StatusOK, u)
}

type updateAllocation struct {
	createRequest
	ID int `json:"ID" binding:"required"`
}

func (u updateAllocation) toAllocation() allocation.Allocation {
	return allocation.Allocation{
		ID:        u.ID,
		UserId:    u.UserId,
		Name:      u.Name,
		Share:     u.Share,
		IsActive:  u.IsActive,
		UpdatedAt: time.Now(),
	}
}

type updateRequest struct {
	Data []updateAllocation `json:"data" binding:"required"`
}

func (er *updateRequest) toAllocations() []allocation.Allocation {
	a := make([]allocation.Allocation, len(er.Data))
	for _, alloc := range er.Data {
		a = append(a, alloc.toAllocation())
	}
	return a
}

func handleUpdate(c *gin.Context) {
	var request updateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	d, err := db.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updated, err := allocation.BulkUpdate(d, request.toAllocations())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}
