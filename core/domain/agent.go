package domain

type Agent struct {
	ID              int64  `gorm:"primaryKey" json:"id"`
	Name            string `gorm:"type varchar(50); not null" json:"name"`
	RequestRecordId int64  `gorm:"not null" json:"request_record_id"`
}
