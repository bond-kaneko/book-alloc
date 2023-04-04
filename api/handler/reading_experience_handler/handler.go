package reading_experience_handler

import (
	"book-alloc/db"
	"book-alloc/internal/reading_experience"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Routes(r *gin.RouterGroup) {
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

	mine := reading_experience.GetMine(d, userId)
	c.JSON(200, mine)
}

type CreateRequest struct {
	AllocationId int    `json:"AllocationId"`
	Title        string `json:"Title"`
	Status       int    `json:"Status"`
}

func (c CreateRequest) toReadingExperience() (reading_experience.ReadingExperience, error) {
	return reading_experience.ReadingExperience{
		AllocationId: c.AllocationId,
		Title:        c.Title,
		Status:       reading_experience.Status(c.Status),
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

	err = reading_experience.Delete(d, reId)
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

func (c UpdateReadingExperience) toReadingExperience() reading_experience.ReadingExperience {
	return reading_experience.ReadingExperience{
		ID:           c.ID,
		AllocationId: c.AllocationId,
		Title:        c.Title,
		Status:       reading_experience.Status(c.Status),
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

	var forUpdate []reading_experience.ReadingExperience
	for _, r := range request {
		forUpdate = append(forUpdate, r.toReadingExperience())
	}

	exps, err := reading_experience.BulkUpdate(d, forUpdate)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, exps)
}
