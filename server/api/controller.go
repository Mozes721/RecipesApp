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
	router.GET("/set-cache", func(c *gin.Context) {
		getCache(c, client)
	})

	router.POST("/get-cache", func(c *gin.Context) {
		setCache(c, client)
	})
}

func SetRoutes(router *gin.Engine, client *firestore.Client, auth *auth.Client) {
	router.OPTIONS("/*any", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.Use(func(c *gin.Context) {
		md.AuthJWT(auth)(c)
	})

	router.GET("/:id", func(c *gin.Context) {
		userID := c.Param("id")

		showRecepies(c, client, userID)
	})

	router.POST("/recipe", func(c *gin.Context) {
		addRecepie(c, client)
	})

	router.PATCH("/recipe/:id", func(c *gin.Context) {
		updateRecepie(c, client)
	})
	router.DELETE("/recipes/:id", func(c *gin.Context) {
		deleteRecepie(c, client)
	})

	router.Run()
}
