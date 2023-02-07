package mocks

// import (
// 	"github.com/justahmed99/delos-farm/core/domain"
// 	"github.com/justahmed99/delos-farm/core/port"
// 	"github.com/stretchr/testify/mock"
// )

// // PondRepository is a mock implementation of the PondRepository interface
// type PondRepository struct {
// 	mock.Mock
// }

// // GetPonds returns the mock implementation of GetPonds
// func (repo *PondRepository) GetPonds() ([]*domain.Pond, error) {
// 	args := repo.Called()
// 	return args.Get(0).([]*domain.Pond), args.Error(1)
// }

// // CreatePond returns the mock implementation of CreatePond
// func (repo *PondRepository) CreatePond(pond *domain.Pond) error {
// 	args := repo.Called(pond)
// 	return args.Error(0)
// }

// // GetPondByID returns the mock implementation of GetPondByID
// func (repo *PondRepository) GetPondByID(id int64) (*domain.Pond, error) {
// 	args := repo.Called(id)
// 	return args.Get(0).(*domain.Pond), args.Error(1)
// }

// // UpdatePond returns the mock implementation of UpdatePond
// func (repo *PondRepository) UpdatePond(pond *domain.Pond) (*domain.Pond, error) {
// 	args := repo.Called(pond)
// 	return args.Get(0).(*domain.Pond), args.Error(1)
// }

// // DeletePond returns the mock implementation of DeletePond
// func (repo *PondRepository) SoftDeletePond(id int64) error {
// 	args := repo.Called(id)
// 	return args.Error(0)
// }

// // Implement the PondRepository interface
// var _ port.PondRepository = (*PondRepository)(nil)
