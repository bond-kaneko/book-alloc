package reading_experience

import (
	"book-alloc/internal/allocation"
	"github.com/pkg/errors"
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

func FromString(value string) (Status, error) {
	switch value {
	case "Want":
		return Want, nil
	case "Reading":
		return Reading, nil
	case "Complete":
		return Complete, nil
	case "Stash":
		return Stash, nil
	}
	return 0, errors.New("Non-existent status")
}

type ReadingExperience struct {
	ID           int
	AllocationId int
	Title        string
	Status       Status
	StartAt      *time.Time
	EndAt        *time.Time
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}

func GetMine(db *gorm.DB, userId string) []ReadingExperience {
	var allocations []allocation.Allocation
	_ = db.Where("user_id = ?", userId).Find(&allocations)

	var allocationIds []int
	for _, a := range allocations {
		allocationIds = append(allocationIds, a.ID)
	}

	var books []ReadingExperience
	_ = db.Where("allocation_id in ?", allocationIds).Find(&books)

	return books
}

func Create(db *gorm.DB, book ReadingExperience) (ReadingExperience, error) {
	err := db.Create(&book).Error
	return book, err
}

func DeleteByAllocationId(db *gorm.DB, allocationId int) error {
	return db.Where("allocation_id = ?", allocationId).Delete(&ReadingExperience{}).Error
}

func Delete(db *gorm.DB, readingExperienceId int) error {
	return db.Where("id =?", readingExperienceId).Delete(&ReadingExperience{}).Error
}

func BulkUpdate(db *gorm.DB, exps []ReadingExperience) ([]ReadingExperience, error) {
	err := db.Transaction(func(tx *gorm.DB) error {
		for _, e := range exps {
			result := tx.Select("*").Omit("id").Where("id = ?", e.ID).Updates(&e)
			if result.Error != nil {
				return result.Error
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return exps, nil
}
