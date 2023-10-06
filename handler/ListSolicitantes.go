package handler

import (
	"net/http"

	"github.com/DaviidSantos/uno-api/schemas"
	"github.com/gin-gonic/gin"
)

func ListSolicitantes(ctx *gin.Context) {
	solicitantes := []schemas.Solicitante{}

	if err := db.Find(&solicitantes).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing solicitantes")
		return
	}

	sendSuccess(ctx, "list-solicitantes", solicitantes)
}