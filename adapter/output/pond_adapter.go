package adapter

import (
	"errors"

	"github.com/justahmed99/delos-farm/core/domain"
	"gorm.io/gorm"
)

type GormPondRepository struct {
	db *gorm.DB
}

// CREATE
func (repo *GormPondRepository) CreatePond(pond *domain.Pond) (*domain.Pond, error) {
	farmRepo := NewGormFarmRepository(repo.db)
	_, err_farm := farmRepo.GetFarmByID(pond.FarmID)
	if err_farm != nil {
		return nil, errors.New("farm not found")
	}
	pond.NewPond(pond.Name, pond.FarmID)
	err := repo.db.Create(pond).Error
	if err != nil {
		return nil, err
	}
	return pond, nil
}

// READ ONE
func (repo *GormPondRepository) GetPondByID(id int64) (*domain.Pond, error) {
	pond := &domain.Pond{}
	err := repo.db.Where("id = ? AND is_active = ?", id, true).First(pond).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return pond, nil
}

// READ ALL
func (repo *GormPondRepository) GetPonds() ([]*domain.Pond, error) {
	var ponds []*domain.Pond
	err := repo.db.Where("is_active = ?", true).Find(&ponds).Error
	if err != nil {
		return nil, err
	}
	return ponds, nil
}

// READ BY FARM ID
func (repo *GormPondRepository) GetPondsByFarmID(farm_id int64) ([]*domain.Pond, error) {
	var ponds []*domain.Pond
	err := repo.db.Where("farm_id = ? AND is_active = ?", farm_id, true).Find(&ponds).Error
	if err != nil {
		return nil, err
	}
	return ponds, nil
}

// UPDATE
func (repo *GormPondRepository) UpdatePond(pond *domain.Pond) (*domain.Pond, error) {
	farmRepo := NewGormFarmRepository(repo.db)
	_, err_farm := farmRepo.GetFarmByID(pond.FarmID)
	if err_farm != nil {
		return nil, errors.New("farm not found")
	}
	affected_row := repo.db.Model(&pond).Where("id = ? AND is_active = ?", pond.ID, true).Updates(&pond).RowsAffected
	if affected_row == 0 {
		insert_new_farm, _ := repo.CreatePond(pond)
		return insert_new_farm, errors.New("Update failed, insert new instead!")
	}
	return pond, nil
}

// DELETE
func (repo *GormPondRepository) SoftDeletePond(id int64) error {
	pond, _ := repo.GetPondByID(id)

	pond.DeletePond()
	err_delete := repo.db.Save(pond).Error
	if err_delete != nil {
		return err_delete
	}
	return nil
}

// DELETE BY FARM ID
func (repo *GormPondRepository) SoftDeletePondsByFarmID(farm_id int64) error {
	ponds, err_find_ponds := repo.GetPondsByFarmID(farm_id)
	if err_find_ponds != nil {
		return err_find_ponds
	}

	if len(ponds) == 0 {
		return nil
	}

	for i := range ponds {
		err := repo.SoftDeletePond(ponds[i].ID)
		if err != nil {
			return errors.New("Error deleting pond!")
		}
	}
	return nil
}
