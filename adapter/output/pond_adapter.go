package adapter

import (
	"github.com/justahmed99/delos-farm/core/domain"
	"gorm.io/gorm"
)

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
