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

const (
	authorizationHeader = "Authorization"
	valName             = "FIREBASE_ID_TOKEN"
)

func AuthJWT(client *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		authHeader := c.Request.Header.Get(authorizationHeader)
		log.Println("authHeader", authHeader)
		token := strings.Replace(authHeader, "Bearer ", "", 1)
		idToken, err := client.VerifyIDToken(c, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		log.Println("Auth time:", time.Since(startTime))

		c.Set(valName, idToken)
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
