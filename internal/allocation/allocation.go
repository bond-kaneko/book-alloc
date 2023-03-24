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

func Create(db *gorm.DB, allocation Allocation) error {
	return db.Omit("ID").Create(&allocation).Error
}

func GetLatestByUserId(d *gorm.DB, userId string) (a Allocation, exists bool) {
	result := d.Where("user_id = ?", userId).Find(&a).Limit(1)
	return a, result.RowsAffected > 0
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

func BulkDelete(db *gorm.DB, ids []int) error {
	return db.Where("id in (?)", ids).Delete(&Allocation{}).Error
}
