package repositories

import (
	"coding_challenge/model"
	"coding_challenge/utils"
	"fmt"

	"gorm.io/gorm"
)

type CarsRepository interface {
	CreateCars(carsPayload *model.Cars) error
	ReadAllCars() ([]*model.Cars, error)
	ReadByIDCars(id int) ([]*model.Cars, error)
	UpdateCars(carsPayload *model.Cars) error
	DeleteCars(id int) error
}

type carsRepository struct {
	db *gorm.DB
}

func (carsRepo *carsRepository) CreateCars(carsPayload *model.Cars) error {
	fmt.Println(carsPayload)
	if err := carsRepo.db.Create(carsPayload).Error; err != nil {
		return &utils.AppError{
			ErrorCode:    201,
			ErrorMessage: err.Error(),
		}
	}
	return nil
}

func (carsRepo *carsRepository) ReadAllCars() ([]*model.Cars, error) {
	var result []*model.Cars
	if err := carsRepo.db.Order("car_id").Find(&result).Error; err != nil {
		return nil, &utils.AppError{
			ErrorCode:    202,
			ErrorMessage: err.Error(),
		}
	}

	fmt.Println(result)
	return result, nil
}

func (carsRepo *carsRepository) ReadByIDCars(id int) ([]*model.Cars, error) {
	fmt.Println(id)
	var result []*model.Cars
	if err := carsRepo.db.Where("car_id = ?", id).Find(&result).Error; err != nil {
		return nil, &utils.AppError{
			ErrorCode:    203,
			ErrorMessage: err.Error(),
		}
	}
	fmt.Println(result)
	return result, nil
}

func (carsRepo *carsRepository) UpdateCars(carsPayload *model.Cars) error {
	fmt.Println(carsPayload)

	query := `UPDATE cars SET car_name = ?, day_rate = ?, month_rate = ?, image = ? WHERE car_id = ?`
	result := carsRepo.db.Exec(query, carsPayload.CarName, carsPayload.DayRate, carsPayload.MonthRate, carsPayload.Image, carsPayload.CarID)

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

func (carsRepo *carsRepository) DeleteCars(id int) error {
	if err := carsRepo.db.Where("car_id = ?", id).Delete(&model.Cars{}).Error; err != nil {
		return &utils.AppError{
			ErrorCode:    205,
			ErrorMessage: err.Error(),
		}
	}
	return nil
}

func NewCarsRepo(db *gorm.DB) CarsRepository {
	return &carsRepository{
		db: db,
	}
}
