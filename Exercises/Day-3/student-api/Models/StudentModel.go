package Models

type Student struct {
	ID      uint   `json:"id"`
	//gorm.Model
	FirstName    string `json:"first_name"`
	LastName   string `json:"last_name"`
	DOB   string `json:"dob"`// time_format:"01-02-2006"`
	Address string `json:"address"`
}
func (b *Student) TableName() string {
	return "students"
}
