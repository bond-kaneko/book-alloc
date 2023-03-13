package user

import (
	"book-alloc/db"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type User struct {
	ID         string
	Name       string
	Email      string
	RegisterAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
}

func GetAll(c *gin.Context) {
	db, _ := db.NewDB()
	var users []User
	_ = db.Find(&users)
	c.JSON(http.StatusOK, users)
}
