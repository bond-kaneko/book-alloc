package reading_history

import (
	"book-alloc/db"
	"time"
)

type ReadingHistory struct {
	ID           int
	AllocationId int
	Isbn         *string
	Title        *string
	Status       int
	Times        int
	StartAt      time.Time
	EndAt        db.NullableTime
	Rating       *int
	Comment      *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func GetAll() []ReadingHistory {
	db, _ := db.NewDB()
	var readingHistories []ReadingHistory
	_ = db.Find(&readingHistories)
	return readingHistories
}
