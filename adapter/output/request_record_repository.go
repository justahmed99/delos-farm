package adapter

import (
	"github.com/justahmed99/delos-farm/core/domain"
	"gorm.io/gorm"
)

type GormRecordRequestRepository struct {
	db *gorm.DB
}

func (repo *GormRecordRequestRepository) SaveRecord(record *domain.RequestRecord) (*domain.RequestRecord, error) {
	err := repo.db.Create(record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (repo *GormRecordRequestRepository) UpdateRecord(record *domain.RequestRecord) (*domain.RequestRecord, error) {
	affected_row := repo.db.Model(&record).Where("record_title = ?", record.RecordTitle).Updates(&record).RowsAffected
	if affected_row == 0 {
		insert_new_farm, _ := repo.SaveRecord(record)
		return insert_new_farm, nil
	}
	return record, nil
}

func (repo *GormRecordRequestRepository) GetRequestRecordById(id int64) (*domain.RequestRecord, error) {
	var agents *domain.RequestRecord
	err := repo.db.Where("record_request_id = ?", id).First(&agents).Error
	if err != nil {
		return nil, err
	}
	return agents, nil
}

func (repo *GormRecordRequestRepository) GetRequestRecordByTitle(title string) (*domain.RequestRecord, error) {
	var agents *domain.RequestRecord
	err := repo.db.Where("record_title = ?", title).First(&agents).Error
	if err != nil {
		return nil, err
	}
	return agents, nil
}

func (repo *GormRecordRequestRepository) GetAllRequestRecords() ([]*domain.RequestRecord, error) {
	var records []*domain.RequestRecord
	err := repo.db.Find(&records).Error
	if err != nil {
		return nil, err
	}
	return records, nil
}
