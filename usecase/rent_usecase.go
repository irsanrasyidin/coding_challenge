package usecase

import (
	"coding_challenge/model"
	"coding_challenge/repositories"
)

type RentUsecase interface {
	CreateRent(rentPayload *model.Orders) error
	ReadAllRent() ([]*model.Orders, error)
	ReadByIDRent(id int) ([]*model.Orders, error)
	UpdateRent(rentPayload *model.Orders) error
	DeleteRent(id int) error
}

type rentUsecase struct {
	rentRepo repositories.RentRepository
}

func (rentUsecase *rentUsecase) CreateRent(rentPayload *model.Orders) error {
	return rentUsecase.rentRepo.CreateRent(rentPayload)
}

func (rentUsecase *rentUsecase) ReadAllRent() ([]*model.Orders, error) {
	return rentUsecase.rentRepo.ReadAllRent()
}

func (rentUsecase *rentUsecase) ReadByIDRent(id int) ([]*model.Orders, error) {
	return rentUsecase.rentRepo.ReadByIDRent(id)
}

func (rentUsecase *rentUsecase) UpdateRent(rentPayload *model.Orders) error {
	return rentUsecase.rentRepo.UpdateRent(rentPayload)
}

func (rentUsecase *rentUsecase) DeleteRent(id int) error {
	return rentUsecase.rentRepo.DeleteRent(id)
}

func NewRentUsecase(rentRepo repositories.RentRepository) RentUsecase {
	return &rentUsecase{
		rentRepo: rentRepo,
	}
}
