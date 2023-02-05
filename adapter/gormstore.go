package adapter

import (
	"github.com/justahmed99/delos-farm/core/port"
	"gorm.io/gorm"
)

func NewGormFarmRepository(db *gorm.DB) port.FarmRepository {
	return &GormFarmRepository{db: db}
}

func NewGormPondRepository(db *gorm.DB) port.PondRepository {
	return &GormFarmRepository{db: db}
}
