package Repo

import (
	"Github-Projects/freshers-bootcamp/Exercises/Day-4/retailer-api/Models"
	"errors"
	"gorm.io/gorm"
)

func CheckCustomerExists(id string, db *gorm.DB) (bool, error) {
	var order Models.Order
	query := db.Select("orders.*")
	query = query.Group("orders.id")
	err := query.Where("customer_id = ?", id).First(&order).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, nil
}
