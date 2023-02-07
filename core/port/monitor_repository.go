package port

import (
	"github.com/gin-gonic/gin"
	"github.com/justahmed99/delos-farm/core/domain"
)

type MonitorRepository interface {
	GetMonitorData(context *gin.Context) (map[string]*domain.Monitor, error)
}
