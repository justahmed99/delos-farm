package adapter

import (
	"github.com/justahmed99/delos-farm/core/port"
	"gorm.io/gorm"
)

type GormAdapter struct {
	db *gorm.DB
}

func NewGormAdapter(db *gorm.DB) *GormAdapter {
	return &GormAdapter{
		db: db,
	}
}

func NewGormFarmRepository(db *gorm.DB) port.FarmRepository {
	return &GormFarmRepository{db: db}
}

func NewGormPondRepository(db *gorm.DB) port.PondRepository {
	return &GormPondRepository{db: db}
}

func NewGormAgentRepository(db *gorm.DB) port.AgentRepository {
	return &GormAgentRepository{db: db}
}

func NewGormRecordRequestRepository(db *gorm.DB) port.RecordRequestRepository {
	return &GormRecordRequestRepository{db: db}
}

func NewGormMonitorRepository(db *gorm.DB) port.MonitorRepository {
	return &GormMonitorRepository{db: db}
}
