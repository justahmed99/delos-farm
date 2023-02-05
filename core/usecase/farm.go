package usecase

import (
	"context"

	"github.com/justahmed99/delos-farm/core/domain"
)

type FarmUseCases interface {
	CreateFarm(context context.Context, name string) (*domain.Farm, error)
	GetFarmByID(context context.Context, id int64) (*domain.Farm, error)
	UpdateFarm(context context.Context, id int64, name string) (*domain.Farm, error)
	DeleteFarm(context context.Context, id int64) (*domain.Farm, error)
}
