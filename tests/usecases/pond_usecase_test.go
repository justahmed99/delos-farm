package tests

import (
	"testing"

	"github.com/justahmed99/delos-farm/core/domain"
	"github.com/justahmed99/delos-farm/core/usecase"
	"github.com/justahmed99/delos-farm/mocks"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestPondUseCase_CreatePond(t *testing.T) {
	repo := new(mocks.PondRepository)
	repo.On("CreatePond", &domain.Pond{
		Name:   "New Pond",
		FarmID: "ff3e9f38-4234-4b0a-b2eb-318fa4f61f16",
	}).Return(&domain.Pond{
		ID:       "ff3e9f38-4234-4b0a-b2eb-318fa4f61f22",
		Name:     "New Pond",
		FarmID:   "ff3e9f38-4234-4b0a-b2eb-318fa4f61f16",
		IsActive: true,
	}, nil)
	usecase := usecase.NewPondUseCases(repo)

	_, err := usecase.CreatePond(nil, domain.Pond{
		Name:   "New Pond",
		FarmID: "ff3e9f38-4234-4b0a-b2eb-318fa4f61f16",
	})

	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestPondUseCase_UpdatePond(t *testing.T) {
	repo := new(mocks.PondRepository)
	repo.On("UpdatePond", &domain.Pond{
		ID:       "ff3e9f38-4234-4b0a-b2eb-318fa4f61f16",
		Name:     "Pond 123",
		FarmID:   "ff3e9f38-4234-4b0a-b2eb-318fa4f61f20",
		IsActive: true,
	}).Return(&domain.Pond{
		ID:       "ff3e9f38-4234-4b0a-b2eb-318fa4f61f16",
		Name:     "Pond 123",
		FarmID:   "ff3e9f38-4234-4b0a-b2eb-318fa4f61f20",
		IsActive: true,
	}, nil)
	usecase := usecase.NewPondUseCases(repo)

	_, err := usecase.UpdatePond(nil, domain.Pond{
		ID:       "ff3e9f38-4234-4b0a-b2eb-318fa4f61f16",
		Name:     "Pond 123",
		FarmID:   "ff3e9f38-4234-4b0a-b2eb-318fa4f61f20",
		IsActive: true,
	})

	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestPondUsecase_SoftDeletePondPositiveTest(t *testing.T) {
	repo := new(mocks.PondRepository)
	repo.On("SoftDeletePond", "ff3e9f38-4234-4b0a-b2eb-318fa4f61f16").Return(nil)
	usecase := usecase.NewPondUseCases(repo)
	err := usecase.DeletePond(nil, "ff3e9f38-4234-4b0a-b2eb-318fa4f61f16")
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestPondUsecase_SoftDeletePondNegativeTest(t *testing.T) {
	repo := new(mocks.PondRepository)
	repo.On("SoftDeletePond", "ff3e9f38-4234-4b0a-b2eb-318fa4f61f16").Return(gorm.ErrRecordNotFound)
	usecase := usecase.NewPondUseCases(repo)
	err := usecase.DeletePond(nil, "ff3e9f38-4234-4b0a-b2eb-318fa4f61f16")
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	repo.AssertExpectations(t)
}
