package usecase

import (
	"context"

	"github.com/justahmed99/delos-farm/core/domain"
)

type PondUseCases interface {
	CreatePond(context context.Context, pond domain.Pond) (*domain.Pond, error)
	GetPondByID(context context.Context, id string) (*domain.Pond, error)
	UpdatePond(context context.Context, pond domain.Pond) (*domain.Pond, error)
	DeletePond(context context.Context, id string) (bool, error)
}
