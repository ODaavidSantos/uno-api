package handler

import (
	"net/http"
	"strings"

	"github.com/DaviidSantos/uno-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update solicitante
// @Description Update a solicitante data
// @Tags Solicitante
// @Accept json
// @Produce json
// @Param request body UpdateSolicitanteRequest true "Request body"
// @Param cnpj query string true "Solicitante identification"
// @Success 200 {object} UpdateSolicitanteResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /solicitante [patch]
func UpdateSolicitante(ctx *gin.Context) {
	request := UpdateSolicitanteRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	cnpj := ctx.Query("cnpj")
	if cnpj == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("cnpj", "queryParameter").Error())
		return
	}

	solicitante := schemas.Solicitante{}

	if err := db.First(&solicitante, cnpj).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "solicitante not found")
		return
	}

	if request.NomeFantasia != "" {
		solicitante.NomeFantasia = request.NomeFantasia
	}
	if request.Ativo != nil {
		solicitante.Ativo = *request.Ativo
	}
	if request.Cep != "" {
		solicitante.Cep = request.Cep
	}
	if request.Rua != "" {
		solicitante.Rua = request.Rua
	}
	if request.Cidade != "" {
		solicitante.Cidade = request.Cidade
	}
	if request.Estado != "" {
		solicitante.Estado = request.Estado
	}

	if err := db.Save(&solicitante).Error; err != nil {
		logger.Errorf("error updating solicitante: %v", err.Error())
		if strings.Contains(err.Error(), "solicitantes_nome_fantasia_key") {
			sendError(ctx, http.StatusConflict, "nome_fantasia already registered!")
			return
		}
		
		sendError(ctx, http.StatusInternalServerError, "error updating solicitante")
		return 
	}

	sendSuccess(ctx, "update-solicitante", solicitante)
}
