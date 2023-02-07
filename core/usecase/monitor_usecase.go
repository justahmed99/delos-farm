package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/justahmed99/delos-farm/core/domain"
	"github.com/justahmed99/delos-farm/core/port"
)

type MonitorUseCases interface {
	GetMonitorData(context *gin.Context) (map[string]*domain.Monitor, error)
}

type MonitorUseCasesImpl struct {
	monitorRepository port.MonitorRepository
}

func NewMonitorUseCases(monitorRepository port.MonitorRepository) MonitorUseCases {
	return &MonitorUseCasesImpl{
		monitorRepository: monitorRepository,
	}
}

func (useCase *MonitorUseCasesImpl) GetMonitorData(context *gin.Context) (map[string]*domain.Monitor, error) {
	return useCase.monitorRepository.GetMonitorData(context)
}
