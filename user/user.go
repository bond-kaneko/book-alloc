package user

import (
	"book-alloc/db"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	id   string `db:"id"`
	name string `db:"name"`
}

type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
}

func GetAll(c *gin.Context) {
	db := db.NewDB()
	c.JSON(http.StatusOK, "get all")
}
