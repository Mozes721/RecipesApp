package api

import (
	"github.com/RecepieApp/server/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func setCache(ctx *gin.Context, client *redis.Client) {
	var userCache models.UserCache

	err := models.UnmarshallRequestBodyToAPIData(ctx.Request.Body, &userCache)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Unable to parse data",
		})
	}

	responseErr := userCache.SetUserCache(ctx, client)
	if responseErr == nil {
		ctx.JSON(400, gin.H{
			"message": responseErr.Error(),
		})
		return
	}

}

func getCache(ctx *gin.Context, client *redis.Client) {
	user := models.User{}

	err := models.UnmarshallRequestBodyToAPIData(ctx.Request.Body, &user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Unable to parse data",
		})
		return
	}

	var userCache models.UserCache
	responseErr := user.GetUserCache(ctx, client, &userCache)
	if responseErr != nil {
		ctx.JSON(400, gin.H{
			"message": userCache,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": userCache,
	})

}
