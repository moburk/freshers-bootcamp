package Repo

import (
	"Github-Projects/freshers-bootcamp/Exercises/Day-4/retailer-api/Models"
	"errors"
	"gorm.io/gorm"
)

func PlaceOrder(order *Models.Order, db *gorm.DB) (err error) {
	if err = db.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderByID(id string, db *gorm.DB) (Models.Order, bool, error) {
	order := Models.Order{}
	query := db.Select("orders.*")
	query = query.Group("orders.id")
	err := query.Where("orders.id = ?", id).First(&order).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return order, false, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return order, false, nil
	}
	return order, true, nil
}
