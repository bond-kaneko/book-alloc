package reading_experience

import (
	"book-alloc/db"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Handle(r *gin.RouterGroup) {
	a := r.Group("/reading-experiences")
	{
		a.GET("/:userId", HandleMyBooks)
		a.POST("/", HandleCreate)
		a.DELETE("/:readingExperienceId", HandleDelete)
		a.PUT("/", HandleUpdate)
	}
}

func HandleMyBooks(c *gin.Context) {
	userId := c.Param("userId")
	d, _ := db.NewDB()

	mine := GetMine(d, userId)
	c.JSON(200, mine)
}

type CreateRequest struct {
	AllocationId int    `json:"AllocationId"`
	Title        string `json:"Title"`
	Status       int    `json:"Status"`
}

func (c CreateRequest) toReadingExperience() (ReadingExperience, error) {
	return ReadingExperience{
		AllocationId: c.AllocationId,
		Title:        c.Title,
		Status:       Status(c.Status),
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
	book, err := Create(d, r)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, book)
}

func HandleDelete(c *gin.Context) {
	reId, err := strconv.Atoi(c.Param("readingExperienceId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	d, err := db.NewDB()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = Delete(d, reId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

type UpdateReadingExperience struct {
	CreateRequest
	ID int `json:"id"`
}

func (c UpdateReadingExperience) toReadingExperience() ReadingExperience {
	return ReadingExperience{
		ID:           c.ID,
		AllocationId: c.AllocationId,
		Title:        c.Title,
		Status:       Status(c.Status),
	}
}

func HandleUpdate(c *gin.Context) {
	var request []UpdateReadingExperience
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	d, err := db.NewDB()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var forUpdate []ReadingExperience
	for _, r := range request {
		forUpdate = append(forUpdate, r.toReadingExperience())
	}

	exps, err := BulkUpdate(d, forUpdate)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, exps)
}
