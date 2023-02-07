package port

import "github.com/justahmed99/delos-farm/core/domain"

type RecordRequestRepository interface {
	SaveRecord(record *domain.RequestRecord) (*domain.RequestRecord, error)
	UpdateRecord(record *domain.RequestRecord) (*domain.RequestRecord, error)
	GetRequestRecordById(id int64) (*domain.RequestRecord, error)
	GetRequestRecordByTitle(title string) (*domain.RequestRecord, error)
	GetAllRequestRecords() ([]*domain.RequestRecord, error)
}
