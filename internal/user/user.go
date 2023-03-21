package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         string
	Auth0Id    string
	RegisterAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func GetByAuth0Id(db *gorm.DB, auth0Id string) (User, error) {
	var user User
	result := db.First(&user, "auth0_id = ?", auth0Id)

	return user, result.Error
}
