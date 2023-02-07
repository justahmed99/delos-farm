package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/justahmed99/delos-farm/core/domain"
	"github.com/justahmed99/delos-farm/core/port"
)

type PondUseCases interface {
	CreatePond(context *gin.Context, pond domain.Pond) (*domain.Pond, error)
	GetPondByID(context *gin.Context, id string) (*domain.Pond, error)
	GetPonds(context *gin.Context) ([]*domain.Pond, error)
	GetPondsByFarmID(context *gin.Context, farm_id string) ([]*domain.Pond, error)
	UpdatePond(context *gin.Context, pond domain.Pond) (*domain.Pond, error)
	DeletePond(context *gin.Context, id string) error
	DeletePondsByFarmID(context *gin.Context, id string) error
}

type PondUseCasesImpl struct {
	pondRepository port.PondRepository
}

func NewPondUseCases(pondRepository port.PondRepository) PondUseCases {
	return &PondUseCasesImpl{
		pondRepository: pondRepository,
	}
}

func (useCase *PondUseCasesImpl) CreatePond(context *gin.Context, pond domain.Pond) (*domain.Pond, error) {
	return useCase.pondRepository.CreatePond(context, &pond)
}

func (useCase *PondUseCasesImpl) GetPondByID(context *gin.Context, id string) (*domain.Pond, error) {
	return useCase.pondRepository.GetPondByID(context, id)
}

func (useCase *PondUseCasesImpl) GetPonds(context *gin.Context) ([]*domain.Pond, error) {
	return useCase.pondRepository.GetPonds(context)
}

func (useCase *PondUseCasesImpl) GetPondsByFarmID(context *gin.Context, farm_id string) ([]*domain.Pond, error) {
	return useCase.pondRepository.GetPondsByFarmID(context, farm_id)
}

func (useCase *PondUseCasesImpl) UpdatePond(context *gin.Context, pond domain.Pond) (*domain.Pond, error) {
	return useCase.pondRepository.UpdatePond(context, &pond)
}

func (useCase *PondUseCasesImpl) DeletePond(context *gin.Context, id string) error {
	return useCase.pondRepository.SoftDeletePond(context, id)
}

func (useCase *PondUseCasesImpl) DeletePondsByFarmID(context *gin.Context, id string) error {
	return useCase.pondRepository.SoftDeletePondsByFarmID(context, id)
}
