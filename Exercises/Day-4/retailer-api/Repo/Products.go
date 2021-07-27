package Repo

import (
	"Github-Projects/freshers-bootcamp/Exercises/Day-4/retailer-api/Models"
	"errors"
	"gorm.io/gorm"
)

//AddProduct Adds new product to the database
func AddProduct(product *Models.Product, db *gorm.DB) (err error) {
	if err = db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

//UpdateProductByID updates product in DB
func UpdateProductByID(product *Models.Product, db *gorm.DB, price int, quantity int) error {
	err := db.Model(product).Updates(Models.Product{Price: price, Quantity: quantity}).Error
	if err != nil {
		return err
	}
	return nil
}

//GetProductByID Fetches product from DB using id
func GetProductByID(id string, db *gorm.DB) (Models.Product, bool, error) {
	product := Models.Product{}
	query := db.Select("products.*")
	query = query.Group("products.id")
	err := query.Where("products.id = ?", id).First(&product).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return product, false, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return product, false, nil
	}
	return product, true, nil
}

func GetProducts(db *gorm.DB) ([]Models.Product, error) {
	var products []Models.Product
	//query := db.Select("products.*").Group("products.id")
	//if err := query.Find(&books).Error
	if err := db.Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}


