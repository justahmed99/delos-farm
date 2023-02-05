package port

import (
	"github.com/justahmed99/delos-farm/core/domain"
)

type FarmRepository interface {
	CreateFarm(farm *domain.Farm) error
	GetFarmByID(id int64) (*domain.Farm, error)
	UpdateFarm(farm *domain.Farm) error
	SoftDeleteFarm(id int64) error
}
