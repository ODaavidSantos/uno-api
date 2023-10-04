package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initilizeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/solicitante", func(ctx * gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Solicitante",
			})
		})

		v1.POST("/solicitante", func(ctx * gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "POST Solicitante",
			})
		})

		v1.PATCH("/solicitante", func(ctx * gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "PATCH Solicitante",
			})
		})

		v1.GET("/solicitantes", func(ctx * gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Get Solicitantes",
			})
		})
	}
}