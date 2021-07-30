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

//GetAllOrders get all orders from database
func GetAllOrders(db *gorm.DB) ([]Models.Order, error) {
	var orders []Models.Order
	if err := db.Find(&orders).Error; err != nil {
		return orders, err
	}
	return orders, nil
}

func GetCustomerOrders(id string, db *gorm.DB) ([]Models.Order, bool, error) {
	var orders []Models.Order
	query := db.Select("orders.*")
	query = query.Group("orders.id")
	err := query.Where("orders.customer_id = ?", id).Find(&orders).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return orders, false, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return orders, false, nil
	}
	return orders, true, nil
}
