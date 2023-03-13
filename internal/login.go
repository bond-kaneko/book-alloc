package internal

import (
	"book-alloc/db"
	"book-alloc/internal/user"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type EmailLoginRequest struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

func Login(c *gin.Context) {
	db, _ := db.NewDB()

	var request EmailLoginRequest
	err := c.BindJSON(&request)
	if err != nil {
		logrus.Info("invalid request: ?", request)
		c.Status(http.StatusBadRequest)
		return
	}

	var user user.User
	result := db.Where("email = ?", request.Email).Find(&user)
	if result.Error != nil {
		logrus.Info("login user not found: ?", request)
		c.Status(http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		logrus.Info("password mismatch")
		c.Status(http.StatusBadRequest)
		return
	}

	session := sessions.Default(c)
	loginUser, err := json.Marshal(user)
	if err != nil {
		logrus.Error("failed user marshaling: ?", user)
		c.Status(http.StatusInternalServerError)
		return
	}

	session.Set("loginUser", string(loginUser))
	session.Save()
	c.Status(http.StatusOK)

}
