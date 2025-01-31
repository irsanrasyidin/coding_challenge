package manager

import (
	"coding_challenge/repositories"
	"sync"
)

type RepoManager interface {
	GetRentRepo() repositories.RentRepository
	GetCarsRepo() repositories.CarsRepository
}

type repomanager struct {
	infraManager InfraManager

	rentRepo repositories.RentRepository
	carsRepo repositories.CarsRepository
}

var onceLoadRentRepo sync.Once
var onceLoadCarsRepo sync.Once

func (rm *repomanager) GetRentRepo() repositories.RentRepository {
	onceLoadRentRepo.Do(func() {
		rm.rentRepo = repositories.NewRentRepo(rm.infraManager.GetDB())
	})

	return rm.rentRepo
}

func (rm *repomanager) GetCarsRepo() repositories.CarsRepository {
	onceLoadCarsRepo.Do(func() {
		rm.carsRepo = repositories.NewCarsRepo(rm.infraManager.GetDB())
	})

	return rm.carsRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repomanager{
		infraManager: infraManager,
	}
}
