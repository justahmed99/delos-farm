package domain

import "errors"

type Pond struct {
	ID        int64
	name      string
	pond_id   int64
	is_active bool
}

func NewPond(name string, pond_id int64) (*Pond, error) {
	if name == "" {
		return nil, errors.New("Name is required")
	}

	if pond_id == 0 {
		return nil, errors.New("Pond ID is required")
	}

	return &Pond{
		name:      name,
		pond_id:   pond_id,
		is_active: true,
	}, nil
}

func (pond *Pond) UpdatePond(name string, pond_id int64) error {
	if name == "" {
		return errors.New("Name is required")
	}

	if pond_id == 0 {
		return errors.New("Pond ID is required")
	}

	pond.name = name
	pond.pond_id = pond_id
	return nil
}

func (pond *Pond) DeletePond() error {
	pond.is_active = false
	return nil
}
