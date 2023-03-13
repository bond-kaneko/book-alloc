package allocation

import (
	"book-alloc/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Allocation struct {
	ID        int
	UserId    string
	Name      string
	Share     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetAll(c *gin.Context) {
	db, _ := db.NewDB()
	var allocations []Allocation
	_ = db.Find(&allocations)
	c.JSON(http.StatusOK, allocations)
}
