package allocation

import (
	"book-alloc/db"
	"book-alloc/internal/reading_experience"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Handle(r *gin.RouterGroup) {
	a := r.Group("/allocations")
	{
		a.POST("/", handleCreate)
		a.PUT("/", handleUpdate)
		a.GET("/:userId", handleGetByUserId)
		a.DELETE("/:allocationId", handleDelete)
		a.GET("/:userId/summary", HandleSummary)
	}
}

type createRequest struct {
	UserId   string `json:"UserId" binding:"required"`
	Name     string `json:"Name" binding:"required"`
	Share    int    `json:"Share" binding:"required"`
	IsActive bool   `json:"IsActive"`
}

func (cr *createRequest) toAllocation() Allocation {
	return Allocation{
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

	alloc, err := Create(d, request.toAllocation())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, alloc)
}

func handleGetByUserId(c *gin.Context) {
	d, err := db.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userId := c.Param("userId")
	u := GetByUserId(d, userId)

	c.JSON(http.StatusOK, u)
}

type updateAllocation struct {
	createRequest
	ID int `json:"ID" binding:"required"`
}

func (u updateAllocation) toAllocation() Allocation {
	return Allocation{
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

	var forUpdate []Allocation
	for _, a := range request {
		forUpdate = append(forUpdate, a.toAllocation())
	}

	updated, err := BulkUpdate(d, forUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func handleDelete(c *gin.Context) {
	d, err := db.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	allocationId, err := strconv.Atoi(c.Param("allocationId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := d.Begin()
	err = reading_experience.DeleteByAllocationId(tx, allocationId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = Delete(tx, allocationId)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func HandleSummary(c *gin.Context) {
	d, err := db.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userId := c.Param("userId")
	countForAllocationId, err := GetReadingExperienceCountForEachAllocation(d, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, countForAllocationId)
}
