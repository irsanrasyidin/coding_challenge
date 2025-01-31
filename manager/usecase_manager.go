package manager

import (
	"coding_challenge/usecase"
	"sync"
)

type UsecaseManager interface {
	GetRentUsecase() usecase.RentUsecase
	GetCarsUsecase() usecase.CarsUsecase
}

type usecasemanager struct {
	repoManager RepoManager

	rentUsecase usecase.RentUsecase
	carsUsecase usecase.CarsUsecase
}

var onceLoadRentUsecase sync.Once
var onceLoadCarsUsecase sync.Once

func (um *usecasemanager) GetRentUsecase() usecase.RentUsecase {
	onceLoadRentUsecase.Do(func() {
		um.rentUsecase = usecase.NewRentUsecase(um.repoManager.GetRentRepo())
	})
	return um.rentUsecase
}

func (um *usecasemanager) GetCarsUsecase() usecase.CarsUsecase {
	onceLoadCarsUsecase.Do(func() {
		um.carsUsecase = usecase.NewCarsUsecase(um.repoManager.GetCarsRepo())
	})
	return um.carsUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecasemanager{
		repoManager: repoManager,
	}
}
