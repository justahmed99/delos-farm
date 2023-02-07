package usecase

import (
	"github.com/justahmed99/delos-farm/core/domain"
	"github.com/justahmed99/delos-farm/core/port"
)

type FarmUseCases interface {
	CreateFarm(farm *domain.Farm) (*domain.Farm, error)
	GetFarmByID(id int64) (*domain.Farm, error)
	GetFarms() ([]*domain.Farm, error)
	UpdateFarm(farm *domain.Farm) (*domain.Farm, error)
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

func (useCase *FarmUseCasesImpl) CreateFarm(farm *domain.Farm) (*domain.Farm, error) {
	return useCase.farmRepository.CreateFarm(farm)
}

func (useCase *FarmUseCasesImpl) GetFarmByID(id int64) (*domain.Farm, error) {
	return useCase.farmRepository.GetFarmByID(id)
}

func (useCase *FarmUseCasesImpl) GetFarms() ([]*domain.Farm, error) {
	return useCase.farmRepository.GetFarms()
}

func (useCase *FarmUseCasesImpl) UpdateFarm(farm *domain.Farm) (*domain.Farm, error) {
	return useCase.farmRepository.UpdateFarm(farm)
}

func (useCase *FarmUseCasesImpl) DeleteFarm(id int64) error {
	return useCase.farmRepository.SoftDeleteFarm(id)
}
