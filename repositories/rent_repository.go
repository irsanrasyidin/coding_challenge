package repositories

import (
	"coding_challenge/model"
	"coding_challenge/utils"
	"fmt"

	"gorm.io/gorm"
)

type RentRepository interface {
	CreateRent(rentPayload *model.Orders) error
	ReadAllRent() ([]*model.Orders, error)
	ReadByIDRent(id int) ([]*model.Orders, error)
	UpdateRent(rentPayload *model.Orders) error
	DeleteRent(id int) error
}

type rentRepository struct {
	db *gorm.DB
}

func (rentRepo *rentRepository) CreateRent(rentPayload *model.Orders) error {
	if err := rentRepo.db.Create(&rentPayload).Error; err != nil {
		return &utils.AppError{
			ErrorCode:    101,
			ErrorMessage: err.Error(),
		}
	}
	return nil
}

func (rentRepo *rentRepository) ReadAllRent() ([]*model.Orders, error) {
	var result []*model.Orders
	if err := rentRepo.db.Order("CAST(order_id AS INTEGER)").Find(&result).Error; err != nil {
		return nil, &utils.AppError{
			ErrorCode:    102,
			ErrorMessage: err.Error(),
		}
	}
	return result, nil
}

func (rentRepo *rentRepository) ReadByIDRent(id int) ([]*model.Orders, error) {
	var result []*model.Orders
	if err := rentRepo.db.Where("order_id = ?", id).Find(&result).Error; err != nil {
		return nil, &utils.AppError{
			ErrorCode:    103,
			ErrorMessage: err.Error(),
		}
	}
	return result, nil
}

func (rentRepo *rentRepository) UpdateRent(rentPayload *model.Orders) error {
	fmt.Println(rentPayload)
	query := `UPDATE orders SET dropoff_location = ? WHERE order_id = ?;`
	result := rentRepo.db.Exec(query, rentPayload.DropoffLocation, rentPayload.CarID)

	if result.RowsAffected == 0 {
		return &utils.AppError{
			ErrorCode:    204,
			ErrorMessage: "No rows were updated, check if car_id exists.",
		}
	}

	if err := result.Error; err != nil {
		return &utils.AppError{
			ErrorCode:    204,
			ErrorMessage: err.Error(),
		}
	}
	return nil
}

func (rentRepo *rentRepository) DeleteRent(id int) error {
	if err := rentRepo.db.Where("order_id = ?", id).Delete(&model.Orders{}).Error; err != nil {
		return &utils.AppError{
			ErrorCode:    105,
			ErrorMessage: err.Error(),
		}
	}
	return nil
}

func NewRentRepo(db *gorm.DB) RentRepository {
	return &rentRepository{
		db: db,
	}
}
