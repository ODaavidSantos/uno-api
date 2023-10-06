package handler

import (
	"net/http"

	"github.com/DaviidSantos/uno-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show solicitante
// @Description Show a solicitante
// @Tags Solicitante
// @Produce json
// @Param cnpj query string true "Solicitante identification"
// @Success 200 {object} ShowSolicitanteResponse
// @Failure 404 {object} ErrorResponse
// @Router /solicitante [get]
func ShowSolicitante(ctx *gin.Context) {
	cnpj := ctx.Query("cnpj")

	if cnpj == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	solicitante := schemas.Solicitante{}
	if err := db.First(&solicitante, cnpj).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "solicitante not found")
		return
	}

	sendSuccess(ctx, "show-solicitante", solicitante)
}