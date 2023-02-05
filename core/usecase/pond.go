package usecase

import (
	"context"

	"github.com/justahmed99/delos-farm/core/domain"
)

type PondUseCases interface {
	CreatePond(context context.Context, name string) (*domain.Pond, error)
	GetPondByID(context context.Context, id int64) (*domain.Pond, error)
	UpdatePond(context context.Context, id int64, name string) (*domain.Pond, error)
	DeletePond(context context.Context, id int64) (*domain.Pond, error)
}
