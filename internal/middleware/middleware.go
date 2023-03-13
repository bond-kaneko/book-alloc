package middleware

import (
	"book-alloc/internal/user"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/koron/go-dproxy"
	"net/http"
)

func LoginCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		loginUserJson, err := dproxy.New(session.Get("loginUser")).String()

		if err != nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
		} else {
			var user user.User
			// Json文字列のアンマーシャル
			err := json.Unmarshal([]byte(loginUserJson), &user)
			if err != nil {
				c.Status(http.StatusUnauthorized)
				c.Abort()
			} else {
				c.Next()
			}
		}
	}
}
