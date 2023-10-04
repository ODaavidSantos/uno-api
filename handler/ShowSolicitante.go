package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowSolicitante(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Get Solicitante",
	})
}