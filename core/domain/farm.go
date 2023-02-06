package domain

import "errors"

type Farm struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type varchar(50); not null" json:"name"`
	IsActive bool   `json:"isActive; not null"`
}

func NewFarm(name string) (*Farm, error) {
	if name == "" {
		return nil, errors.New("Name is required")
	}

	return &Farm{
		Name:     name,
		IsActive: true,
	}, nil
}

func (farm *Farm) UpdateFarm(name string) error {
	if name == "" {
		return errors.New("Name is required")
	}
	farm.Name = name
	return nil
}

func (farm *Farm) DeleteFarm() error {
	farm.IsActive = false
	return nil
}
