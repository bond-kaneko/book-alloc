package user

import (
	"book-alloc/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type User struct {
	ID         string
	Name       string
	Email      string
	Password   string
	RegisterAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func GetAll(c *gin.Context) {
	db, _ := db.NewDB()
	var users []User
	_ = db.Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetByEmail(c *gin.Context, email string) (User, error) {
	db, _ := db.NewDB()

	var user User
	result := db.First(&user, "email = ?", email)

	return user, result.Error
}
