package mocks

import (
	"github.com/gin-gonic/gin"
	"github.com/justahmed99/delos-farm/core/domain"
	"github.com/justahmed99/delos-farm/core/port"
	"github.com/stretchr/testify/mock"
)

type FarmRepository struct {
	mock.Mock
}

func (repo *FarmRepository) GetFarms(context *gin.Context) ([]*domain.Farm, error) {
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

func (repo *FarmRepository) CreateFarm(context *gin.Context, farm *domain.Farm) (*domain.Farm, error) {
	args := repo.Called(farm)
	return args.Get(0).(*domain.Farm), args.Error(1)
}

func (repo *FarmRepository) GetFarmByID(context *gin.Context, id string) (*domain.Farm, error) {
	args := repo.Called(id)
	return args.Get(0).(*domain.Farm), args.Error(1)
}

func (repo *FarmRepository) UpdateFarm(context *gin.Context, farm *domain.Farm) (*domain.Farm, error) {
	args := repo.Called(farm)
	return args.Get(0).(*domain.Farm), args.Error(1)
}

func (repo *FarmRepository) SoftDeleteFarm(context *gin.Context, id string) error {
	args := repo.Called(id)
	return args.Error(0)
}

var _ port.FarmRepository = (*FarmRepository)(nil)
