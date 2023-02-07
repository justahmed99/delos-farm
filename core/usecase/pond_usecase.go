package usecase

import (
	"github.com/justahmed99/delos-farm/core/domain"
	"github.com/justahmed99/delos-farm/core/port"
)

type PondUseCases interface {
	CreatePond(pond domain.Pond) (*domain.Pond, error)
	GetPondByID(id int64) (*domain.Pond, error)
	GetPonds() ([]*domain.Pond, error)
	GetPondsByFarmID(farm_id int64) ([]*domain.Pond, error)
	UpdatePond(pond domain.Pond) (*domain.Pond, error)
	DeletePond(id int64) error
	DeletePondsByFarmID(id int64) error
}

type PondUseCasesImpl struct {
	pondRepository port.PondRepository
}

func NewPondUseCases(pondRepository port.PondRepository) PondUseCases {
	return &PondUseCasesImpl{
		pondRepository: pondRepository,
	}
}

func (useCase *PondUseCasesImpl) CreatePond(pond domain.Pond) (*domain.Pond, error) {
	return useCase.pondRepository.CreatePond(&pond)
}

func (useCase *PondUseCasesImpl) GetPondByID(id int64) (*domain.Pond, error) {
	return useCase.pondRepository.GetPondByID(id)
}

func (useCase *PondUseCasesImpl) GetPonds() ([]*domain.Pond, error) {
	return useCase.pondRepository.GetPonds()
}

func (useCase *PondUseCasesImpl) GetPondsByFarmID(farm_id int64) ([]*domain.Pond, error) {
	return useCase.pondRepository.GetPondsByFarmID(farm_id)
}

func (useCase *PondUseCasesImpl) UpdatePond(pond domain.Pond) (*domain.Pond, error) {
	return useCase.pondRepository.UpdatePond(&pond)
}

func (useCase *PondUseCasesImpl) DeletePond(id int64) error {
	return useCase.pondRepository.SoftDeletePond(id)
}

func (useCase *PondUseCasesImpl) DeletePondsByFarmID(id int64) error {
	return useCase.pondRepository.SoftDeletePondsByFarmID(id)
}
