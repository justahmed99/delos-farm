package adapter

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/justahmed99/delos-farm/core/domain"
	"github.com/justahmed99/delos-farm/program_constant"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type GormPondRepository struct {
	db *gorm.DB
}

// CREATE
func (repo *GormPondRepository) CreatePond(context *gin.Context, pond *domain.Pond) (*domain.Pond, error) {
	if context != nil {
		_, err_record := UpdateRequestRecordAndAgentForPond(repo, context, "POST /pond")
		if err_record != nil {
			return nil, err_record
		}
	}

	farmRepo := NewGormFarmRepository(repo.db)
	_, err_farm := farmRepo.GetFarmByID(nil, pond.FarmID)
	if err_farm != nil {
		return nil, errors.New(program_constant.FARM_NOT_FOUND_ERROR)
	}
	pond.ID = uuid.NewV4().String()
	pond.NewPond(pond.Name, pond.FarmID)
	err := repo.db.Create(pond).Error
	if err != nil {
		return nil, err
	}
	return pond, nil
}

// READ ONE
func (repo *GormPondRepository) GetPondByID(context *gin.Context, id string) (*domain.Pond, error) {
	if context != nil {
		_, err_record := UpdateRequestRecordAndAgentForPond(repo, context, "GET /pond/:id")
		if err_record != nil {
			return nil, err_record
		}
	}

	pond := &domain.Pond{}
	err := repo.db.Where("id = ? AND is_active = ?", id, true).First(pond).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return pond, nil
}

// READ ALL
func (repo *GormPondRepository) GetPonds(context *gin.Context) ([]*domain.Pond, error) {
	if context != nil {
		_, err_record := UpdateRequestRecordAndAgentForPond(repo, context, "GET /pond")
		if err_record != nil {
			return nil, err_record
		}
	}

	var ponds []*domain.Pond
	err := repo.db.Where("is_active = ?", true).Find(&ponds).Error
	if err != nil {
		return nil, err
	}
	return ponds, nil
}

// READ BY FARM ID
func (repo *GormPondRepository) GetPondsByFarmID(context *gin.Context, farm_id string) ([]*domain.Pond, error) {
	if context != nil {
		_, err_record := UpdateRequestRecordAndAgentForPond(repo, context, "GET /pond/farm/:id")
		if err_record != nil {
			return nil, err_record
		}
	}

	var ponds []*domain.Pond
	err := repo.db.Where("farm_id = ? AND is_active = ?", farm_id, true).Find(&ponds).Error
	if err != nil {
		return nil, err
	}
	return ponds, nil
}

// UPDATE
func (repo *GormPondRepository) UpdatePond(context *gin.Context, pond *domain.Pond) (*domain.Pond, error) {
	if context != nil {
		_, err_record := UpdateRequestRecordAndAgentForPond(repo, context, "PUT /pond")
		if err_record != nil {
			return nil, err_record
		}
	}

	farmRepo := NewGormFarmRepository(repo.db)
	_, err_farm := farmRepo.GetFarmByID(nil, pond.FarmID)
	if err_farm != nil {
		return nil, errors.New("farm not found")
	}
	affected_row := repo.db.Model(&pond).Where("id = ? AND is_active = ?", pond.ID, true).Updates(&pond).RowsAffected
	if affected_row == 0 {
		insert_new_farm, _ := repo.CreatePond(nil, pond)
		return insert_new_farm, errors.New(program_constant.UPDATE_ERROR)
	}
	return pond, nil
}

// DELETE
func (repo *GormPondRepository) SoftDeletePond(context *gin.Context, id string) error {
	if context != nil {
		_, err_record := UpdateRequestRecordAndAgentForPond(repo, context, "DELETE /pond")
		if err_record != nil {
			return err_record
		}
	}

	pond, _ := repo.GetPondByID(nil, id)

	pond.IsActive = false
	err_delete := repo.db.Save(pond).Error
	if err_delete != nil {
		return err_delete
	}
	return nil
}

// DELETE BY FARM ID
func (repo *GormPondRepository) SoftDeletePondsByFarmID(context *gin.Context, farm_id string) error {
	if context != nil {
		_, err_record := UpdateRequestRecordAndAgentForPond(repo, context, "PUT /pond/farm/:id")
		if err_record != nil {
			return err_record
		}
	}

	ponds, err_find_ponds := repo.GetPondsByFarmID(nil, farm_id)
	if err_find_ponds != nil {
		return err_find_ponds
	}

	if len(ponds) == 0 {
		return nil
	}

	for i := range ponds {
		err := repo.SoftDeletePond(nil, ponds[i].ID)
		if err != nil {
			return errors.New(program_constant.POND_DELETE_ERROR)
		}
	}
	return nil
}
