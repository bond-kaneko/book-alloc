package allocation

import (
	"book-alloc/db"
	"book-alloc/internal/allocation"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
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
	UserId   string
	Name     string
	Share    int
	IsActive bool
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

type updateRequest struct {
	Allocations []allocation.Allocation `json:"data"`
}

func (er *updateRequest) toAllocations() []allocation.Allocation {
	fmt.Println(er.Allocations)
	a := make([]allocation.Allocation, len(er.Allocations))
	for _, alloc := range er.Allocations {
		alloc.UpdatedAt = time.Now()
		a = append(a, alloc)
	}
	return a
}

func handleUpdate(c *gin.Context) {
	dump, _ := httputil.DumpRequest(c.Request, true)
	fmt.Println(string(dump))

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
