package usecase

import (
	"coding_challenge/model"
	"coding_challenge/repositories"
	"fmt"
)

type CarsUsecase interface {
	CreateCars(carsPayload *model.Cars) error
	ReadAllCars() ([]*model.Cars, error)
	ReadByIDCars(id int) ([]*model.Cars, error)
	UpdateCars(carsPayload *model.Cars) error
	DeleteCars(id int) error
}

type carsUsecase struct {
	carsRepo repositories.CarsRepository
}

func (carsUsecase *carsUsecase) CreateCars(carsPayload *model.Cars) error {
	fmt.Println("usecase=", carsPayload)
	return carsUsecase.carsRepo.CreateCars(carsPayload)
}

func (carsUsecase *carsUsecase) ReadAllCars() ([]*model.Cars, error) {
	return carsUsecase.carsRepo.ReadAllCars()
}

func (carsUsecase *carsUsecase) ReadByIDCars(id int) ([]*model.Cars, error) {
	return carsUsecase.carsRepo.ReadByIDCars(id)
}

func (carsUsecase *carsUsecase) UpdateCars(carsPayload *model.Cars) error {
	return carsUsecase.carsRepo.UpdateCars(carsPayload)
}

func (carsUsecase *carsUsecase) DeleteCars(id int) error {
	return carsUsecase.carsRepo.DeleteCars(id)
}

func NewCarsUsecase(carsRepo repositories.CarsRepository) CarsUsecase {
	return &carsUsecase{
		carsRepo: carsRepo,
	}
}
