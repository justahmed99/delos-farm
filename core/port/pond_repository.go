package port

import "github.com/justahmed99/delos-farm/core/domain"

type PondRepository interface {
	CreatePond(pond *domain.Pond) (*domain.Pond, error)
	GetPondByID(id int64) (*domain.Pond, error)
	GetPonds() ([]*domain.Pond, error)
	GetPondsByFarmID(farm_id int64) ([]*domain.Pond, error)
	UpdatePond(pond *domain.Pond) (*domain.Pond, error)
	SoftDeletePond(id int64) error
	SoftDeletePondsByFarmID(farm_id int64) error
}
