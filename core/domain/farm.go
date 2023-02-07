package domain

import (
	"time"
)

type Farm struct {
	ID        string     `gorm:"type:varchar(36);primaryKey" json:"id"`
	Name      string     `gorm:"type varchar(50); not null" json:"name"`
	IsActive  bool       `gorm:"not null; default:true" json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `gorm:"default:null" json:"updated_at"`
}
