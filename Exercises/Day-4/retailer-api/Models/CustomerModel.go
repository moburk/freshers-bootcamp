package Models

type Customer struct {
	//internalID uint `gorm:"primaryKey"`
	ID string `json:"id" gorm:"size:10"`
	CustomerName string `json:"product_name"`
}

func (c *Customer) TableName() string {
	return "customers"
}
