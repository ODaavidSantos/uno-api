package handler

import (
	"net/http"

	"github.com/DaviidSantos/uno-api/schemas"
	"github.com/gin-gonic/gin"
)

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