package tests

import (
	"testing"

	"github.com/justahmed99/delos-farm/core/domain"
	"github.com/justahmed99/delos-farm/core/usecase"
	"github.com/justahmed99/delos-farm/mocks"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestFarmUseCase_GetFarmsAllIsActive(t *testing.T) {
	repo := new(mocks.FarmRepository)
	repo.On("GetFarms").Return([]*domain.Farm{
		{ID: 1, Name: "Farm 1", IsActive: true},
		{ID: 2, Name: "Farm 2", IsActive: true},
		{ID: 2, Name: "Farm 3", IsActive: true},
		{ID: 2, Name: "Farm 4", IsActive: true},
	}, nil)
	usecase := usecase.NewFarmUsesCases(repo)
	farms, err := usecase.GetFarms()
	assert.NoError(t, err)
	assert.Len(t, farms, 4)

	repo.AssertExpectations(t)
}

func TestFarmUseCase_GetFarmsNotAllIsActive(t *testing.T) {
	repo := new(mocks.FarmRepository)
	repo.On("GetFarms").Return([]*domain.Farm{
		{ID: 1, Name: "Farm 1", IsActive: false},
		{ID: 2, Name: "Farm 2", IsActive: true},
		{ID: 3, Name: "Farm 3", IsActive: false},
		{ID: 4, Name: "Farm 4", IsActive: true},
	}, nil)
	usecase := usecase.NewFarmUsesCases(repo)
	farms, err := usecase.GetFarms()
	assert.NoError(t, err)
	assert.Len(t, farms, 2)

	repo.AssertExpectations(t)
}

func TestFarmUseCase_CreateFarm(t *testing.T) {
	repo := new(mocks.FarmRepository)
	repo.On("CreateFarm", &domain.Farm{
		Name: "New Farm",
	}).Return(&domain.Farm{
		ID:       1,
		Name:     "New Farm",
		IsActive: true,
	}, nil)
	usecase := usecase.NewFarmUsesCases(repo)

	_, err := usecase.CreateFarm(&domain.Farm{
		Name: "New Farm",
	})

	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestFarmUseCase_UpdateFarm(t *testing.T) {
	repo := new(mocks.FarmRepository)
	repo.On("UpdateFarm", &domain.Farm{
		ID:       9,
		Name:     "Life Farm",
		IsActive: true,
	}).Return(&domain.Farm{
		ID:       9,
		Name:     "Life Farm",
		IsActive: true,
	}, nil)
	usecase := usecase.NewFarmUsesCases(repo)

	_, err := usecase.UpdateFarm(&domain.Farm{
		ID:       9,
		Name:     "Life Farm",
		IsActive: true,
	})

	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestFarmUsecase_SoftDeleteFarmPositiveTest(t *testing.T) {
	repo := new(mocks.FarmRepository)
	repo.On("SoftDeleteFarm", int64(1)).Return(nil)
	usecase := usecase.NewFarmUsesCases(repo)
	err := usecase.DeleteFarm(int64(1))
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestFarmUsecase_SoftDeleteFarmNegativeTest(t *testing.T) {
	repo := new(mocks.FarmRepository)
	repo.On("SoftDeleteFarm", int64(2)).Return(gorm.ErrRecordNotFound)
	usecase := usecase.NewFarmUsesCases(repo)
	err := usecase.DeleteFarm(int64(2))
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	repo.AssertExpectations(t)
}
