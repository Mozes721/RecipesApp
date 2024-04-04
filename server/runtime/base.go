package runtime

import (
	"github.com/RecepieApp/server/api"
	"github.com/RecepieApp/server/app"
	md "github.com/RecepieApp/server/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start(a *app.Application) error {
	router := gin.New()

	router.Use(cors.New(md.CORSMiddleware()))

	api.SetCache(router, a.RedisClient)

	api.SetRoutes(router, a.FireClient, a.FireAuth, a.RedisClient)

	err := router.Run(":" + a.ListenPort)
	if err != nil {
		return err
	}

	return nil
}
