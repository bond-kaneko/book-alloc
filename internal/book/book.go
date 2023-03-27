package book

import (
	"book-alloc/internal/allocation"
	"gorm.io/gorm"
	"time"
)

type Status int

const (
	Want Status = iota
	Reading
	Complete
	Stash
)

type Book struct {
	ID           int
	AllocationId int
	Title        string
	Status       Status
	StartAt      *time.Time
	EndAt        *time.Time
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}

func GetMyBooks(db *gorm.DB, userId string) []Book {
	var allocations []allocation.Allocation
	_ = db.Where("user_id = ?", userId).Find(&allocations)

	var allocationIds []int
	for _, a := range allocations {
		allocationIds = append(allocationIds, a.ID)
	}

	var books []Book
	_ = db.Where("allocation_id in ?", allocationIds).Find(&books)

	return books
}
