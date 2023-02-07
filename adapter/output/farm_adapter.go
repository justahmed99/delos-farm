package adapter

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/justahmed99/delos-farm/core/domain"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type GormFarmRepository struct {
	db *gorm.DB
}

// CREATE
func (repo *GormFarmRepository) CreateFarm(context *gin.Context, farm *domain.Farm) (*domain.Farm, error) {
	if context != nil {
		_, err_record := UpdateRequestRecordAndAgentForFarm(repo, context, "POST /farm")
		if err_record != nil {
			return nil, err_record
		}
	}

	farm.ID = uuid.NewV4().String()
	farm.IsActive = true
	err := repo.db.Create(farm).Error
	if err != nil {
		return nil, err
	}
	return farm, nil
}

// READ ONE
func (repo *GormFarmRepository) GetFarmByID(context *gin.Context, id string) (*domain.Farm, error) {
	if context != nil {
		_, err_record := UpdateRequestRecordAndAgentForFarm(repo, context, "GET /farm/:id")
		if err_record != nil {
			return nil, err_record
		}
	}

	farm := &domain.Farm{}
	err := repo.db.Where("id = ? AND is_active = ?", id, true).First(farm).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	return farm, nil
}

// READ ALL
func (repo *GormFarmRepository) GetFarms(context *gin.Context) ([]*domain.Farm, error) {
	if context != nil {
		_, err_record := UpdateRequestRecordAndAgentForFarm(repo, context, "GET /farm")
		if err_record != nil {
			return nil, err_record
		}
	}

	var farms []*domain.Farm
	err := repo.db.Where("is_active = ?", true).Find(&farms).Error
	if err != nil {
		return nil, err
	}
	return farms, nil
}

// UPDATE
func (repo *GormFarmRepository) UpdateFarm(context *gin.Context, farm *domain.Farm) (*domain.Farm, error) {
	if context != nil {
		_, err_record := UpdateRequestRecordAndAgentForFarm(repo, context, "PUT /farm")
		if err_record != nil {
			return nil, err_record
		}
	}

	farm.IsActive = true
	farm.UpdatedAt = &time.Time{}
	affected_row := repo.db.Model(&farm).Where("id = ? AND is_active = ?", farm.ID, true).Updates(&farm).RowsAffected
	if affected_row == 0 {
		insert_new_farm, _ := repo.CreateFarm(nil, farm)
		return insert_new_farm, errors.New("update failed, insert new instead!")
	}
	return farm, nil
}

// DELETE
func (repo *GormFarmRepository) SoftDeleteFarm(context *gin.Context, id string) error {
	if context != nil {
		_, err_record := UpdateRequestRecordAndAgentForFarm(repo, context, "DELETE /farm")
		if err_record != nil {
			return err_record
		}
	}

	farm, farm_err := repo.GetFarmByID(nil, id)
	if farm_err != nil {
		return farm_err
	}

	pondRepo := NewGormPondRepository(repo.db)
	err_delete_ponds := pondRepo.SoftDeletePondsByFarmID(nil, id)
	if err_delete_ponds != nil {
		return err_delete_ponds
	}

	farm.IsActive = false
	err_delete := repo.db.Save(farm).Error
	if err_delete != nil {
		return err_delete
	}
	return nil
}
