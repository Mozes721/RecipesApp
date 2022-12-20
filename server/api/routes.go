package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	cors := func(c *gin.Context) {
		c.Writer.Header().Add("access-control-allow-origin", "*")
		c.Writer.Header().Add("access-control-allow-headers", "accept, content-type")
		c.Writer.Header().Add("access-control-allow-methods", "GET,HEAD,POST,DELETE,OPTIONS,PUT,PATCH")
	}
	r.Use(cors)
	recepiesRouter := r.Group("/v1")
	{
		recepiesRouter.GET("/", allRecepies)
		//recepiesRouter.POST("/new-recepie/:id", NewRecepie)
		//recepiesRouter.DELETE("/delete/:id", DeleteRecepie)
		//recepiesRouter.PATCH("/find-recepie/:id", UpdateRecepie)
	}
}
