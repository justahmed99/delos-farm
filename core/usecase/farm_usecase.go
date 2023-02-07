package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/justahmed99/delos-farm/core/domain"
	"github.com/justahmed99/delos-farm/core/port"
)

type FarmUseCases interface {
	CreateFarm(context *gin.Context, farm *domain.Farm) (*domain.Farm, error)
	GetFarmByID(context *gin.Context, id string) (*domain.Farm, error)
	GetFarms(context *gin.Context) ([]*domain.Farm, error)
	UpdateFarm(context *gin.Context, farm *domain.Farm) (*domain.Farm, error)
	DeleteFarm(context *gin.Context, id string) error
}

type FarmUseCasesImpl struct {
	farmRepository port.FarmRepository
}

func NewFarmUsesCases(farmRepository port.FarmRepository) FarmUseCases {
	return &FarmUseCasesImpl{
		farmRepository: farmRepository,
	}
}

func (useCase *FarmUseCasesImpl) CreateFarm(context *gin.Context, farm *domain.Farm) (*domain.Farm, error) {
	return useCase.farmRepository.CreateFarm(context, farm)
}

func (useCase *FarmUseCasesImpl) GetFarmByID(context *gin.Context, id string) (*domain.Farm, error) {
	return useCase.farmRepository.GetFarmByID(context, id)
}

func (useCase *FarmUseCasesImpl) GetFarms(context *gin.Context) ([]*domain.Farm, error) {
	return useCase.farmRepository.GetFarms(context)
}

func (useCase *FarmUseCasesImpl) UpdateFarm(context *gin.Context, farm *domain.Farm) (*domain.Farm, error) {
	return useCase.farmRepository.UpdateFarm(context, farm)
}

func (useCase *FarmUseCasesImpl) DeleteFarm(context *gin.Context, id string) error {
	return useCase.farmRepository.SoftDeleteFarm(context, id)
}
