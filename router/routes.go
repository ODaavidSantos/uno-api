package router

import (
	"github.com/DaviidSantos/uno-api/handler"
	"github.com/gin-gonic/gin"
)

func initilizeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/solicitante", handler.ShowSolicitante)

		v1.POST("/solicitante", handler.CreateSolicitante)

		v1.PATCH("/solicitante", handler.UpdateSolicitante)

		v1.GET("/solicitantes", handler.ListSolicitantes)
	}
}