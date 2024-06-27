package api

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	md "github.com/RecepieApp/server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
)

func SetCache(router *gin.Engine, client *redis.Client) {
	router.POST("/set-cache", func(c *gin.Context) {
		setUserCache(c, client)
	})

	router.GET("/check-expiration", func(c *gin.Context) {
		checkTokenExpiration(c, client)
	})

}

func SetRoutes(router *gin.Engine, client *firestore.Client, auth *auth.Client, redisClient *redis.Client) {
	router.OPTIONS("/*any", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.Use(func(c *gin.Context) {
		authToken := getUserCache(c, redisClient)
		md.AuthJWT(auth, authToken)(c)
	})

	router.GET("/", func(c *gin.Context) {
		showRecepies(c, client)
	})

	router.POST("/", func(c *gin.Context) {
		addRecepie(c, client)
	})

	router.PATCH("/", func(c *gin.Context) {
		updateRecepie(c, client)
	})

	router.DELETE("/:id", func(c *gin.Context) {
		deleteRecepie(c, client)
	})

	router.Run()
}
