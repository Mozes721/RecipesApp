package middleware

import (
	"firebase.google.com/go/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"
	"log"
	"net/http"
	"strings"
	"time"
)

func AuthJWT(client *auth.Client, authToken string) gin.HandlerFunc {

	return func(c *gin.Context) {
		token := strings.Replace(authToken, "Bearer ", "", 1)

		idToken, err := client.VerifyIDTokenAndCheckRevoked(c, token)
		if err != nil {
			log.Printf("Token verification failed: %v", err)

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Set("FIREBASE_ID_TOKEN", idToken.UID)
		c.Next()
	}
}

func CORSMiddleware() cors.Config {
	clientPort := envy.Get("REACT_PORT", "http://localhost:3000")

	corsConfig := cors.Config{
		AllowOrigins:     []string{clientPort},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	return corsConfig
}
