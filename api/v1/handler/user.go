package handler

import (
	"book-alloc/internal/user"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func LoginUserHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"user": claims[user.IdentityKey],
	})
}
