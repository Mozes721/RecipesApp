package api

import (
	"context"
	"fmt"
	"github.com/RecepieApp/server/app"
	"github.com/RecepieApp/server/db"
	"github.com/gin-gonic/gin"
)

func allRecepies(c *gin.Context) {
	ctx := context.Background()
	client, err := app.LoadFirebaseConnection()
	defer client.Close()
	if err != nil {
		err = fmt.Errorf("error occurred while loading")
	}
	ok, _ := db.ReadCollection(ctx, client)

	c.JSON(200, ok)
}
