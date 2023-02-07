package domain

type RequestRecord struct {
	ID              int64  `gorm:"primaryKey" json:"id"`
	RecordTitle     string `gorm:"type varchar(50); not null; unique" json:"record_title"`
	Count           int64  `gorm:"not null; default:0" json:"count"`
	UniqueUserAgent int64  `gorm:"not null; default:0" json:"unique_user_agent"`
}
