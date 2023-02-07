package dto

import (
	"time"

	"github.com/justahmed99/delos-farm/core/domain"
)

type FarmDto struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`
}

type FarmRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type FarmResponse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	IsActive  bool       `json:"isActive"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func ConvertFarmRequestToFarm(dto *FarmRequest) *domain.Farm {
	return &domain.Farm{
		Name: dto.Name,
	}
}
