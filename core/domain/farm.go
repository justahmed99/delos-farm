package domain

import "errors"

type Farm struct {
	ID        int64
	name      string
	is_active bool
}

func NewFarm(name string) (*Farm, error) {
	if name == "" {
		return nil, errors.New("Name is required")
	}

	return &Farm{
		name:      name,
		is_active: true,
	}, nil
}

func (farm *Farm) UpdateFarm(name string) error {
	if name == "" {
		return errors.New("Name is required")
	}
	farm.name = name
	return nil
}

func (farm *Farm) DeleteFarm() error {
	farm.is_active = false
	return nil
}
