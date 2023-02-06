package adapter

import (
	"errors"

	"github.com/justahmed99/delos-farm/core/domain"
	"gorm.io/gorm"
)

type GormAdapter struct {
	db *gorm.DB
}

func NewGormAdapter(db *gorm.DB) *GormAdapter {
	return &GormAdapter{
		db: db,
	}
}

// ////// FARM REPOSITORY START
type GormFarmRepository struct {
	db *gorm.DB
}

func (r *GormFarmRepository) CreateFarm(farm *domain.Farm) (*domain.Farm, error) {
	farm.NewFarm(farm.Name)
	err := r.db.Create(farm).Error
	if err != nil {
		return nil, err
	}
	return farm, nil
}

func (r *GormFarmRepository) GetFarmByID(id int64) (*domain.Farm, error) {
	farm := &domain.Farm{}
	err := r.db.Where("id = ? AND is_active = ?", id, true).First(farm).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	return farm, nil
}

func (r *GormFarmRepository) GetFarms() ([]*domain.Farm, error) {
	var farms []*domain.Farm
	err := r.db.Where("is_active = ?", true).Find(&farms).Error
	if err != nil {
		return nil, err
	}
	return farms, nil
}

func (r *GormFarmRepository) UpdateFarm(farm *domain.Farm) (*domain.Farm, error) {
	farm.UpdateFarm(farm.Name)
	affected_row := r.db.Model(&farm).Where("id = ? AND is_active = ?", farm.ID, true).Updates(&farm).RowsAffected
	if affected_row == 0 {
		insert_new_farm, _ := r.CreateFarm(farm)
		return insert_new_farm, errors.New("Update failed, insert new instead!")
	}
	return farm, nil
}

func (r *GormFarmRepository) SoftDeleteFarm(id int64) error {
	farm := &domain.Farm{}
	err_find := r.db.First(farm, id).Error
	if err_find == gorm.ErrRecordNotFound {
		return err_find
	}

	farm.DeleteFarm()
	err_delete := r.db.Save(farm).Error
	if err_delete != nil {
		return err_delete
	}
	return nil
}

//////// FARM REPOSITORY END

//////// POND REPOSITORY START

type GormPondRepository struct {
	db *gorm.DB
}

func (r *GormFarmRepository) CreatePond(pond *domain.Pond) error {
	return r.db.Create(pond).Error
}

func (r *GormFarmRepository) GetPondByID(id int64) (*domain.Pond, error) {
	pond := &domain.Pond{}
	err := r.db.First(pond, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return pond, nil
}

func (r *GormFarmRepository) UpdatePond(pond *domain.Pond) error {
	return r.db.Save(pond).Error
}

func (r *GormFarmRepository) SoftDeletePond(id int64) error {
	pond := &domain.Pond{}
	pond.ID = id
	pond.DeletePond()
	return r.db.Save(pond).Error
}

//////// POND REPOSITORY END
