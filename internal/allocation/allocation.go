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
	result := d.Find(&a).Limit(1)
	return a, result.RowsAffected > 0
}
