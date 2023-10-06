package router

import (
	"github.com/DaviidSantos/uno-api/docs"
	"github.com/DaviidSantos/uno-api/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initilizeRoutes(r *gin.Engine) {
	// Initialize Handler
	handler.Init()
	
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	// Defining routes
	v1 := r.Group(basePath)
	{
		v1.GET("/solicitante", handler.ShowSolicitante)

		v1.POST("/solicitante", handler.CreateSolicitante)

		v1.PATCH("/solicitante", handler.UpdateSolicitante)

		v1.GET("/solicitantes", handler.ListSolicitantes)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}