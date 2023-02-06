package domain

import "errors"

type Farm struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type varchar(50); not null" json:"name"`
	IsActive bool   `gorm:"not null; default:true" json:"isActive"`
}

func (farm *Farm) NewFarm(name string) (*Farm, error) {
	if name == "" {
		return nil, errors.New("Name is required")
	}

	return &Farm{
		Name:     name,
		IsActive: true,
	}, nil
}

func (farm *Farm) UpdateFarm(name string) (*Farm, error) {
	if name == "" {
		return nil, errors.New("Name is required")
	}
	return &Farm{
		Name: name,
	}, nil
	// farm.Name = name
	// return nil
}

func (farm *Farm) DeleteFarm() error {
	farm.IsActive = false
	return nil
}
