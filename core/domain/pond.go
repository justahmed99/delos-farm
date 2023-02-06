package domain

import "errors"

type Pond struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(50);not null" json:"name"`
	FarmID   int64  `gorm:"type:int64;not null" json:"pondId"`
	IsActive bool   `json:"isActive"`
}

func NewPond(name string, farm_id int64) (*Pond, error) {
	if name == "" {
		return nil, errors.New("Name is required")
	}

	if farm_id == 0 {
		return nil, errors.New("Pond ID is required")
	}

	return &Pond{
		Name:     name,
		FarmID:   farm_id,
		IsActive: true,
	}, nil
}

func (pond *Pond) UpdatePond(name string, farm_id int64) error {
	if name == "" {
		return errors.New("Name is required")
	}

	if farm_id == 0 {
		return errors.New("Pond ID is required")
	}

	pond.Name = name
	pond.FarmID = farm_id
	return nil
}

func (pond *Pond) DeletePond() error {
	pond.IsActive = false
	return nil
}
