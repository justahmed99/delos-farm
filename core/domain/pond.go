package domain

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Pond struct {
	ID        string     `gorm:"type:varchar(36);primaryKey" json:"id"`
	Name      string     `gorm:"type:varchar(50);not null" json:"name"`
	FarmID    string     `gorm:"type:varchar(36);index; not null" json:"farm_id"`
	IsActive  bool       `gorm:"default:true; not null" json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `gorm:"default:null" json:"updated_at"`
}

func (pond *Pond) NewPond(name string, farm_id string) (*Pond, error) {
	if name == "" {
		return nil, errors.New("Name is required")
	}

	if farm_id == "" {
		return nil, errors.New("Pond ID is required")
	}

	var id string = uuid.NewV4().String()

	return &Pond{
		ID:       id,
		Name:     name,
		FarmID:   farm_id,
		IsActive: true,
	}, nil
}
