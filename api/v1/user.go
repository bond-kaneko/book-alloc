package v1

import (
	"book-alloc/db"
	"book-alloc/internal/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Auth0IdRequest struct {
	Auth0Id string
}

func User(r *gin.RouterGroup) {
	u := r.Group("/users")
	{
		u.POST("/me", func(c *gin.Context) {
			var request Auth0IdRequest
			if err := c.ShouldBindJSON(&request); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "auth0Id is required"})
			}

			d, err := db.NewDB()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "There is a problem with the database connection"})
			}

			u, err := user.GetByAuth0Id(d, request.Auth0Id)
			switch err {
			case gorm.ErrRecordNotFound:
			// TODO sign up
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user data"})
				return
			}

			c.JSON(http.StatusOK, u)
		})
	}
}
