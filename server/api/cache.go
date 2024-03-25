package api

import (
	"fmt"
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

	key := fmt.Sprintf("user:%s", userCache.UserID)
	val, existsErr := client.HGetAll(ctx, key).Result()

	switch {
	case existsErr == redis.Nil:
		responseErr := userCache.SetUserCache(ctx, client, key)
		if responseErr != nil {
			ctx.JSON(400, gin.H{
				"message": responseErr.Error(),
			})
			return
		}
	case val != nil:
		return
	}

}

func getCache(ctx *gin.Context, client *redis.Client) {
	userID := ctx.Request.URL.Query().Get("userID")

	var userCache models.UserCache
	cache, responseErr := models.GetUserCache(ctx, client, userID)
	if responseErr != nil {
		ctx.JSON(400, gin.H{
			"message": responseErr,
		})
		return
	}
	err := models.UnmarshallCacheData(cache, &userCache)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err,
		})
	}
	ctx.JSON(200, gin.H{
		"message": userCache,
	})

}
