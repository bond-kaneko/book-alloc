package allocation

import (
	"gorm.io/gorm"
	"time"
)

type Allocation struct {
	ID        int
	UserId    string
	Name      string
	Share     int
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Create(db *gorm.DB, allocation Allocation) (Allocation, error) {
	err := db.Create(&allocation).Error
	return allocation, err
}

func GetByUserId(d *gorm.DB, userId string) (a []Allocation) {
	// TODO ページネーション
	// TODO エラーハンドリング
	_ = d.Where("user_id = ?", userId).Find(&a)
	return a
}

func BulkUpdate(db *gorm.DB, allocations []Allocation) ([]Allocation, error) {
	var a []Allocation
	err := db.Transaction(func(tx *gorm.DB) error {
		for _, allocation := range allocations {
			result := tx.Select("*").Omit("id").Model(&a).Where("id = ?", allocation.ID).Updates(&allocation)
			if result.Error != nil {
				return result.Error
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return allocations, nil
}

func Delete(db *gorm.DB, id int) error {
	return db.Where("id = ?", id).Delete(&Allocation{}).Error
}
