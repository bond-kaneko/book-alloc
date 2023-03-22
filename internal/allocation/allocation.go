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
