package runtime

import (
	"github.com/RecepieApp/server/api"
	"github.com/RecepieApp/server/app"
	md "github.com/RecepieApp/server/middleware"
	"log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start(a *app.Application) error {
	router := gin.New()

	router.Use(cors.New(md.CORSMiddleware()))

	api.SetCache(router, a.RedisClient)

	api.SetRoutes(router, a.FireClient, a.FireAuth, a.RedisClient)

	err := router.Run(":" + a.ListenPort)
	log.Printf("Starting server on port: %s", a.ListenPort)
	if err != nil {
		return err
	}

	return nil
}
