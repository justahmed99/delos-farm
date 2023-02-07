package port

import (
	"github.com/gin-gonic/gin"
	"github.com/justahmed99/delos-farm/core/domain"
)

type FarmRepository interface {
	CreateFarm(context *gin.Context, farm *domain.Farm) (*domain.Farm, error)
	GetFarmByID(context *gin.Context, id string) (*domain.Farm, error)
	GetFarms(context *gin.Context) ([]*domain.Farm, error)
	UpdateFarm(context *gin.Context, farm *domain.Farm) (*domain.Farm, error)
	SoftDeleteFarm(context *gin.Context, id string) error
}
