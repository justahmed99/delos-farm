package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justahmed99/delos-farm/core/usecase"
)

type MonitorHandler struct {
	monitorUseCase usecase.MonitorUseCases
}

func NewMonitorHandler(monitorUseCase usecase.MonitorUseCases) *MonitorHandler {
	return &MonitorHandler{
		monitorUseCase: monitorUseCase,
	}
}

func (h *MonitorHandler) GetMonitorData(context *gin.Context) {
	monitor_data, err := h.monitorUseCase.GetMonitorData(context)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	context.IndentedJSON(http.StatusOK, monitor_data)
}
