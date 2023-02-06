package usecase

import (
	"github.com/justahmed99/delos-farm/core/domain"
	"github.com/justahmed99/delos-farm/core/port"
)

type FarmUseCases interface {
	CreateFarm(farm domain.Farm) (*domain.Farm, error)
	GetFarmByID(id int64) (*domain.Farm, error)
	UpdateFarm(farm domain.Farm) (*domain.Farm, error)
	DeleteFarm(id int64) error
}

type FarmUseCasesImpl struct {
	farmRepository port.FarmRepository
}

func NewFarmUsesCases(farmRepository port.FarmRepository) FarmUseCases {
	return &FarmUseCasesImpl{
		farmRepository: farmRepository,
	}
}

func (u *FarmUseCasesImpl) CreateFarm(farm domain.Farm) (*domain.Farm, error) {
	return u.farmRepository.CreateFarm(&farm)
	// return nil, nil
}

func (u *FarmUseCasesImpl) GetFarmByID(id int64) (*domain.Farm, error) {
	return u.farmRepository.GetFarmByID(id)
}

func (u *FarmUseCasesImpl) UpdateFarm(farm domain.Farm) (*domain.Farm, error) {
	return u.farmRepository.UpdateFarm(&farm)
}

func (u *FarmUseCasesImpl) DeleteFarm(id int64) error {
	return u.farmRepository.SoftDeleteFarm(id)
}
