package book

import (
	"book-alloc/db"
	"book-alloc/internal/allocation"
	"github.com/gin-gonic/gin"
	"time"
)

type Status int

const (
	want Status = iota
	reading
	complete
	stash
)

type Book struct {
	ID           int
	AllocationId int
	Title        string
	Status       int
	StartAt      *time.Time
	EndAt        *time.Time
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}

func Handle(r *gin.RouterGroup) {
	a := r.Group("/books")
	{
		a.GET("/:userId", func(c *gin.Context) {
			userId := c.Param("userId")

			d, _ := db.NewDB()

			var allocations []allocation.Allocation
			_ = d.Where("user_id = ?", userId).Find(&allocations)

			var allocationIds []int
			for _, a := range allocations {
				allocationIds = append(allocationIds, a.ID)
			}

			var books []Book
			_ = d.Where("allocation_id in ?", allocationIds).Find(&books)

			c.JSON(200, books)
		})
	}
}
