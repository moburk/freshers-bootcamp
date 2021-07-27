package Models

type Customer struct {
	//internalID uint `gorm:"primaryKey"`
	ID string `json:"id" gorm:"size:10"`
	CustomerName string `json:"product_name"`
	Email string `json:"email"`
}

func (c *Customer) TableName() string {
	return "customers"
}
