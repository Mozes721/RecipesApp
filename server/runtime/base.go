package runtime

import (
	"cloud.google.com/go/firestore"
	"github.com/RecepieApp/server/api"
	"github.com/RecepieApp/server/config"
	"github.com/gin-gonic/gin"
)

func Start(client *firestore.Client) error {
	router := gin.Default()

	router.Use(config.Cors())

	api.SetRoutes(router, client)

	return nil
}
