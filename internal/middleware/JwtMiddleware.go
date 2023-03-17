package middleware

import (
	"book-alloc/internal/user"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const IdentityKey = "email"

type login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func NewJwtMiddleware() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*user.User); ok {
				return jwt.MapClaims{
					IdentityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &user.User{
				Email: claims[IdentityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginParams login
			if err := c.ShouldBind(&loginParams); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			u, err := user.GetByEmail(c, loginParams.Email)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(loginParams.Password))
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return u, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		return nil, err
	}

	return authMiddleware, nil
}
