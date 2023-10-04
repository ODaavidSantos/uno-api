package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListSolicitantes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Get Solicitantes",
	})
}