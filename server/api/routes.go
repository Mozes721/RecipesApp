package api

import "github.com/gin-gonic/gin"

type Handler interface {
	GetAllRecepies(c *gin.Context)
	AddRecepie(c *gin.Context)
	UpdateRecepie(c *gin.Context)
	DeleteRecepie(c *gin.Context)
}

func SetRoutes(engine *gin.Engine, h Handler) {

}
