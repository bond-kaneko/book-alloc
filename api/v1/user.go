package v1

import (
	"book-alloc/db"
	"book-alloc/internal/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IdentifyRequest struct {
	Auth0Id string
	Email   string
	Name    string
}

func User(r *gin.RouterGroup) {
	u := r.Group("/users")
	{
		u.POST("/me", func(c *gin.Context) {
			var request IdentifyRequest
			if err := c.ShouldBindJSON(&request); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "auth0Id is required"})
			}

			d, err := db.NewDB()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "There is a problem with the database connection"})
			}

			u, exists := user.GetByAuth0Id(d, request.Auth0Id)
			if !exists {
				newUser := user.User{
					Auth0Id: request.Auth0Id,
					Email:   request.Email,
					Name:    request.Name,
				}
				err := user.Create(d, newUser)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
					return
				}

				u = newUser
			}

			c.JSON(http.StatusOK, u)
		})
	}
}
