package reading_experience

import (
	"book-alloc/db"
	"book-alloc/internal/reading_experience"
	"github.com/gin-gonic/gin"
	"time"
)

func Handle(r *gin.RouterGroup) {
	a := r.Group("/reading-experiences")
	{
		a.GET("/:userId", HandleMyBooks)
		a.POST("/", HandleCreate)
	}
}

func HandleMyBooks(c *gin.Context) {
	userId := c.Param("userId")
	d, _ := db.NewDB()

	mine := reading_experience.GetMine(d, userId)
	c.JSON(200, mine)
}

type CreateRequest struct {
	AllocationId int    `json:"AllocationId"`
	Title        string `json:"Title"`
	Status       string `json:"Status"`
}

func (c CreateRequest) toReadingExperience() (reading_experience.ReadingExperience, error) {
	s, err := reading_experience.FromString(c.Status)
	if err != nil {
		return reading_experience.ReadingExperience{}, err
	}

	return reading_experience.ReadingExperience{
		AllocationId: c.AllocationId,
		Title:        c.Title,
		Status:       s,
		StartAt:      db.TimePointer(time.Now()),
	}, nil
}

func HandleCreate(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	d, _ := db.NewDB()
	r, err := req.toReadingExperience()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	book, err := reading_experience.Create(d, r)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, book)
}
