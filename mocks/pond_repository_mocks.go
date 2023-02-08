package mocks

import (
	"github.com/gin-gonic/gin"
	"github.com/justahmed99/delos-farm/core/domain"
	"github.com/justahmed99/delos-farm/core/port"
	"github.com/stretchr/testify/mock"
)

type PondRepository struct {
	mock.Mock
}

func (repo *PondRepository) GetPonds(context *gin.Context) ([]*domain.Pond, error) {
	args := repo.Called()
	return args.Get(0).([]*domain.Pond), args.Error(1)
}

func (repo *PondRepository) GetPondsByFarmID(context *gin.Context, farm_id string) ([]*domain.Pond, error) {
	args := repo.Called()
	return args.Get(0).([]*domain.Pond), args.Error(1)
}

func (repo *PondRepository) CreatePond(context *gin.Context, pond *domain.Pond) (*domain.Pond, error) {
	args := repo.Called(pond)
	return args.Get(0).(*domain.Pond), args.Error(1)
}

func (repo *PondRepository) GetPondByID(context *gin.Context, id string) (*domain.Pond, error) {
	args := repo.Called(id)
	return args.Get(0).(*domain.Pond), args.Error(1)
}

func (repo *PondRepository) UpdatePond(context *gin.Context, pond *domain.Pond) (*domain.Pond, error) {
	args := repo.Called(pond)
	return args.Get(0).(*domain.Pond), args.Error(1)
}

func (repo *PondRepository) SoftDeletePond(context *gin.Context, id string) error {
	args := repo.Called(id)
	return args.Error(0)
}

func (repo *PondRepository) SoftDeletePondsByFarmID(context *gin.Context, farm_id string) error {
	args := repo.Called(farm_id)
	return args.Error(0)
}

var _ port.PondRepository = (*PondRepository)(nil)
