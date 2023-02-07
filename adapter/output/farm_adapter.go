package adapter

import (
	"errors"

	"github.com/justahmed99/delos-farm/core/domain"
	"gorm.io/gorm"
)

type GormFarmRepository struct {
	db *gorm.DB
}

// CREATE
func (repo *GormFarmRepository) CreateFarm(farm *domain.Farm) (*domain.Farm, error) {
	farm.NewFarm(farm.Name)
	err := repo.db.Create(farm).Error
	if err != nil {
		return nil, err
	}
	return farm, nil
}

// READ ONE
func (repo *GormFarmRepository) GetFarmByID(id int64) (*domain.Farm, error) {
	farm := &domain.Farm{}
	err := repo.db.Where("id = ? AND is_active = ?", id, true).First(farm).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	return farm, nil
}

// READ ALL
func (repo *GormFarmRepository) GetFarms() ([]*domain.Farm, error) {
	var farms []*domain.Farm
	err := repo.db.Where("is_active = ?", true).Find(&farms).Error
	if err != nil {
		return nil, err
	}
	return farms, nil
}

// UPDATE
func (repo *GormFarmRepository) UpdateFarm(farm *domain.Farm) (*domain.Farm, error) {
	farm.UpdateFarm(farm.Name)
	affected_row := repo.db.Model(&farm).Where("id = ? AND is_active = ?", farm.ID, true).Updates(&farm).RowsAffected
	if affected_row == 0 {
		insert_new_farm, _ := repo.CreateFarm(farm)
		return insert_new_farm, errors.New("Update failed, insert new instead!")
	}
	return farm, nil
}

// DELETE
func (repo *GormFarmRepository) SoftDeleteFarm(id int64) error {
	farm, farm_err := repo.GetFarmByID(id)
	if farm_err != nil {
		return farm_err
	}
	farm.DeleteFarm()
	err_delete := repo.db.Save(farm).Error
	if err_delete != nil {
		return err_delete
	}
	return nil
}
