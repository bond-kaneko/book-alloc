package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           string
	Auth0Id      string
	Email        string
	Name         string
	RegisteredAt time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func GetByAuth0Id(db *gorm.DB, auth0Id string) (user User, exists bool) {
	result := db.Find(&user, "auth0_id = ?", auth0Id).Limit(1)

	return user, result.RowsAffected > 0
}

func Create(db *gorm.DB, user User) error {
	return db.Omit("ID").Create(&user).Error
}
