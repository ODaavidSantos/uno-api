package handler

import (
	"net/http"

	"github.com/DaviidSantos/uno-api/schemas"
	"github.com/gin-gonic/gin"
)


// @BasePath /api/v1

// @Summary List solicitante
// @Description List all solicitantes
// @Tags Solicitante
// @Produce json
// @Success 200 {object} ListSolicitanteResponse
// @Failure 500 {object} ErrorResponse
// @Router /solicitantes [get]
func ListSolicitantes(ctx *gin.Context) {
	solicitantes := []schemas.Solicitante{}

	if err := db.Find(&solicitantes).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing solicitantes")
		return
	}

	sendSuccess(ctx, "list-solicitantes", solicitantes)
}