package main

import (
	"book-alloc/middleware"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	router := http.NewServeMux()

	// This route is always accessible.
	router.Handle("/api/public", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Access-Control-Allow-Origin, Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,UPDATE,OPTIONS")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Hello from a public endpoint! You don't need to be authenticated to see this."}`))
	}))

	// This route is only accessible if the user has a valid access_token.
	router.Handle("/api/private", middleware.EnsureValidToken()(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// CORS Headers.
			//w.Header().Set("Access-Control-Allow-Credentials", "true")
			//w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
			//w.Header().Set("Access-Control-Allow-Headers", "Authorization, Access-Control-Allow-Origin, Content-Type")
			//w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,UPDATE,OPTIONS")
			//w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Hello from a private endpoint! You need to be authenticated to see this."}`))
		}),
	))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	handler := c.Handler(router)
	log.Print("Server listening on http://localhost:8888")
	if err := http.ListenAndServe("0.0.0.0:8080", handler); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}

	//if err := godotenv.Load(); err != nil {
	//	log.Fatalf("Error loading the .env file: %v", err)
	//}
	//
	//g := gin.Default()
	//
	//g.Use(corsMiddleware())
	//
	//g.GET("/ping", func(g *gin.Context) {
	//	g.JSON(200, gin.H{"text": "Hello from public"})
	//})
	//
	//g.GET("/secured/ping", checkJWT(), func(g *gin.Context) {
	//	g.JSON(200, gin.H{"text": "Hello from private"})
	//})

	//g.Use(cors.New(cors.Config{
	//	AllowOrigins: []string{
	//		"*",
	//	},
	//	AllowHeaders: []string{
	//		"Access-Control-Allow-Credentials",
	//		"Access-Control-Allow-Headers",
	//		"Content-Type",
	//		"Content-Length",
	//		"Accept-Encoding",
	//		"Authorization",
	//	},
	//}))

	//authMiddleware, err := middleware.NewJwtMiddleware()
	//if err != nil {
	//	logrus.Error("setUp auth failed: ?", err)
	//}
	//g.POST("/login", authMiddleware.LoginHandler)
	//g.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
	//	claims := jwt.ExtractClaims(c)
	//	log.Printf("NoRoute claims: %#v\n", claims)
	//	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	//})

	//auth := g.Group("/auth")
	//auth.POST("/logout", authMiddleware.LogoutHandler)
	//auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	//auth.Use(authMiddleware.EnsureValidToken())
	//{
	//	route := auth.Group("/v1")
	//	v1.User(route)
	//	v1.Allocation(route)
	//	v1.ReadingHistory(route)
	//}

	//g.Run()
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3333")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func checkJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtMid := *jwtMiddleware
		if err := jwtMid.CheckJWT(c.Writer, c.Request); err != nil {
			c.AbortWithStatus(401)
		}
	}
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("uYtCR0xwYtohm96HaYzYHZ0FCSPdlf4KkjAyoyRNqyGJjXvXlRrL7tWRJ5vuV_iz"), nil
	},
	SigningMethod: jwt.SigningMethodRS256,
})
