package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/justahmed99/delos-farm/core/domain"
	"gorm.io/gorm"
)

type GormMonitorRepository struct {
	db *gorm.DB
}

func (repo *GormMonitorRepository) GetMonitorData(context *gin.Context) (map[string]*domain.Monitor, error) {
	requestRecordRepo := NewGormRecordRequestRepository(repo.db)
	recods, err_recors := requestRecordRepo.GetAllRequestRecords()

	agentRepo := NewGormAgentRepository(repo.db)

	if err_recors != nil {
		return nil, err_recors
	}
	var values = make(map[string]*domain.Monitor)
	for _, record := range recods {
		agent_count, err_agent_count := agentRepo.CountUniqueAgent(record.ID)
		if err_agent_count != nil {
			return nil, err_agent_count
		}
		values[record.RecordTitle] = &domain.Monitor{
			Count:           record.Count,
			UniqueUserAgent: agent_count,
		}
	}

	return values, nil
}
