package adapter

import (
	"github.com/justahmed99/delos-farm/core/domain"
	"gorm.io/gorm"
)

type GormAgentRepository struct {
	db *gorm.DB
}

func (repo *GormAgentRepository) SaveAgent(agent *domain.Agent) (*domain.Agent, error) {
	err := repo.db.Create(agent).Error
	if err != nil {
		return nil, err
	}
	return agent, nil
}

func (repo *GormAgentRepository) GetAgentsByRequestRecordIdAndName(recordId int64, name string) (*domain.Agent, error) {
	agents := &domain.Agent{}
	err := repo.db.Where("name = ? AND request_record_id = ?", name, recordId).First(&agents).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	return agents, nil
}

func (repo *GormAgentRepository) CountUniqueAgent(recordId int64) (int64, error) {
	var count int64
	err := repo.db.Model(&domain.Agent{}).Where("request_record_id = ?", recordId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
