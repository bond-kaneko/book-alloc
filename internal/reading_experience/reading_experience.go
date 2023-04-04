package reading_experience

import (
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
	var aIds []int
	db.Model(&ReadingExperience{}).
		Raw(`
			SELECT id 
			FROM allocations
			WHERE user_id = ?
		`, userId).
		Scan(&aIds)

	var books []ReadingExperience
	_ = db.Where("allocation_id in ?", aIds).Find(&books)

	return books
}

func Create(db *gorm.DB, book ReadingExperience) (ReadingExperience, error) {
	err := db.Create(&book).Error
	return book, err
}

func GetCountForEachAllocationId(db *gorm.DB, allocationIds []int) (map[int]int, error) {
	type result struct {
		AllocationId int
		Count        int
	}
	var res []result
	d := db.Raw(`
		SELECT allocation_id, count(id) as count
		FROM reading_experiences
		WHERE allocation_id IN (?)
		GROUP BY "allocation_id"`, allocationIds).
		Scan(&res)
	if d.Error != nil {
		return nil, d.Error
	}

	countForAllocationId := make(map[int]int)
	for _, r := range res {
		countForAllocationId[r.AllocationId] = r.Count
	}

	return countForAllocationId, nil
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
