package api

import (
	"github.com/gin-gonic/gin"
)

func allRecepies(c *gin.Context) {

	c.JSON(200, "Hello, world!")

}
