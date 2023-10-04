package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateSolicitante(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Patch Solicitante",
	})
}
