package port

import (
	"github.com/gin-gonic/gin"
	"github.com/justahmed99/delos-farm/core/domain"
)

type PondRepository interface {
	CreatePond(context *gin.Context, pond *domain.Pond) (*domain.Pond, error)
	GetPondByID(context *gin.Context, id string) (*domain.Pond, error)
	GetPonds(context *gin.Context) ([]*domain.Pond, error)
	GetPondsByFarmID(context *gin.Context, farm_id string) ([]*domain.Pond, error)
	UpdatePond(context *gin.Context, pond *domain.Pond) (*domain.Pond, error)
	SoftDeletePond(context *gin.Context, id string) error
	SoftDeletePondsByFarmID(context *gin.Context, farm_id string) error
}
