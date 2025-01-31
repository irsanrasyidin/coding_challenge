package model

import "time"

type Cars struct {
	CarID     int     `gorm:"column:car_id;type:integer;primaryKey;autoIncrement" json:"car_id"`
	CarName   string  `gorm:"column:car_name;type:varchar(50);not null" json:"car_name"`
	DayRate   float64 `gorm:"column:day_rate;type:double precision;not null" json:"day_rate"`
	MonthRate float64 `gorm:"column:month_rate;type:double precision;not null" json:"month_rate"`
	Image     string  `gorm:"column:image;type:varchar(256);not null" json:"image"`
}

func (Cars) TableName() string {
	return "cars"
}

type Orders struct {
	OrderID         int       `gorm:"column:order_id;type:integer;primaryKey;autoIncrement" json:"order_id"`
	CarID           int       `gorm:"column:car_id;type:integer;not null" json:"car_id"`
	OrderDate       time.Time `gorm:"column:order_date;type:date;not null" json:"order_date"`
	PickupDate      time.Time `gorm:"column:pickup_date;type:date;not null" json:"pickup_date"`
	DropoffDate     time.Time `gorm:"column:dropoff_date;type:date;not null" json:"dropoff_date"`
	PickupLoaction  string    `gorm:"column:pickup_location;type:varchar(50);not null" json:"pickup_location"`
	DropoffLocation string    `gorm:"column:dropoff_location;type:varchar(50);not null" json:"dropoff_location"`
}

func (Orders) TableName() string {
	return "orders"
}
