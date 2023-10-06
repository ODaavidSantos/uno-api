package handler

import (
	"net/http"

	"github.com/DaviidSantos/uno-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create solicitante
// @Description Create a new solicitante
// @Tags Solicitante
// @Accept json
// @Produce json
// @Param request body CreateSolicitanteRequest true "Request body"
// @Success 200 {object} CreateSolicitanteResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /solicitante [post]
func CreateSolicitante(ctx *gin.Context) {
	request := CreateSolicitanteRequest{}

	ctx.BindJSON(&request)

	logger.Infof("request received: %+v", request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	solicitante := schemas.Solicitante{
		Cnpj: request.Cnpj,
		NomeFantasia: request.NomeFantasia,
		Cep: request.Cep,
		Rua: request.Rua,
		Numero: request.Numero,
		Cidade: request.Cidade,
		Estado: request.Estado,
	}

	if err := db.Create(&solicitante).Error; err != nil {
		logger.Errorf("Error creating solicitante: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create-solicitante", solicitante)
}