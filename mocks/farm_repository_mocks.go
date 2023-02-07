package mocks

import (
	"github.com/justahmed99/delos-farm/core/domain"
	"github.com/justahmed99/delos-farm/core/port"
	"github.com/stretchr/testify/mock"
)

type FarmRepository struct {
	mock.Mock
}

func (repo *FarmRepository) GetFarms() ([]*domain.Farm, error) {
	args := repo.Called()
	farms := args.Get(0).([]*domain.Farm)
	var activeFarms []*domain.Farm
	for _, farm := range farms {
		if farm.IsActive {
			activeFarms = append(activeFarms, farm)
		}
	}
	return activeFarms, args.Error(1)
}

func (repo *FarmRepository) CreateFarm(farm *domain.Farm) (*domain.Farm, error) {
	args := repo.Called(farm)
	return args.Get(0).(*domain.Farm), args.Error(1)
}

func (repo *FarmRepository) GetFarmByID(id int64) (*domain.Farm, error) {
	args := repo.Called(id)
	return args.Get(0).(*domain.Farm), args.Error(1)
}

func (repo *FarmRepository) UpdateFarm(farm *domain.Farm) (*domain.Farm, error) {
	args := repo.Called(farm)
	return args.Get(0).(*domain.Farm), args.Error(1)
}

func (repo *FarmRepository) SoftDeleteFarm(id int64) error {
	args := repo.Called(id)
	return args.Error(0)
}

var _ port.FarmRepository = (*FarmRepository)(nil)
