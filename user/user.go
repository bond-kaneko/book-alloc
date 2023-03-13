package user

import (
	"book-alloc/db"
	"context"
	"fmt"
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
	db, err := db.NewDB()
	fmt.Println(err)
	var u []User
	_ = db.Find(&u)
	c.JSON(http.StatusOK, u)
}
