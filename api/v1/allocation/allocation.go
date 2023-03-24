package allocation

import (
	"book-alloc/db"
	"book-alloc/internal/allocation"
	"fmt"
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
	IsActive bool   `json:"IsActive"`
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
	ID        int  `json:"ID" binding:"required"`
	IsDeleted bool `json:"IsDeleted"`
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

func handleUpdate(c *gin.Context) {
	var request []updateAllocation
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	d, err := db.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var forUpdate []allocation.Allocation
	var forDeleteIds []int
	for _, alloc := range request {
		switch {
		case alloc.IsDeleted:
			forDeleteIds = append(forDeleteIds, alloc.ID)
		default:
			forUpdate = append(forUpdate, alloc.toAllocation())
		}
	}

	updated, err := allocation.BulkUpdate(d, forUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = allocation.BulkDelete(d, forDeleteIds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}
