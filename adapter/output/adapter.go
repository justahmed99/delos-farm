package adapter

import (
	"github.com/justahmed99/delos-farm/core/domain"
	// "github.com/justahmed99/delos-farm/core/port"
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
	farm.IsActive = true
	err := r.db.Create(farm).Error
	if err != nil {
		return nil, err
	}
	return farm, nil
}

func (r *GormFarmRepository) GetFarmByID(id int64) (*domain.Farm, error) {
	farm := &domain.Farm{}
	err := r.db.First(farm, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	return farm, nil
}

func (r *GormFarmRepository) UpdateFarm(farm *domain.Farm) (*domain.Farm, error) {
	err := r.db.Save(farm).Error
	if err != nil {
		return nil, err
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
	// farm.DeleteFarm()
	// err := r.db.Save(farm).Error
	// if err != nil {
	// 	return nil
	// }
	// return err
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
