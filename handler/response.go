package handler

import (
	"fmt"
	"net/http"

	"github.com/DaviidSantos/uno-api/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
	ErrorCode string `json:"error_code"`
}

type CreateSolicitanteResponse struct {
	Message string `json:"message"`
	Data schemas.Solicitante `json:"data"`
}

type ListSolicitanteResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.Solicitante 		`json:"data"`
}

type ShowSolicitanteResponse struct {
	Message string                    `json:"message"`
	Data    schemas.Solicitante 	  `json:"data"`
}

type UpdateSolicitanteResponse struct {
	Message string `json:"message"`
	Data schemas.Solicitante `json:"data"`
}