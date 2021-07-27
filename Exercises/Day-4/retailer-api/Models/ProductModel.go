package Models

type Product struct {
	//internalID uint `gorm:"primaryKey"`
	ID string `json:"id" gorm:"size:10"`
	ProductName string `json:"product_name"`
	Price int `json:"price"`
	Quantity int `json:"quantity"`
}


func (p *Product) TableName() string {
	return "products"
}
