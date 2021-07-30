package Models

import (
	"time"
)

type Order struct {
	ID string `json:"id" gorm:"size:10""`
	CreatedAt time.Time
	CustomerID string `json:"customer_id"`
	customer Customer `gorm:"foreignkey:customer_id; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductID string `json:"product_id"`
	product Product `gorm:"foreignkey:product_id; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Quantity int `json:"quantity"`
	Status string `json:"status"`
}

 func (o *Order) TableName() string {
 	return "orders"
 }