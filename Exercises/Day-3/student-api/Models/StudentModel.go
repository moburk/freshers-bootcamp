package Models

type Student struct {
	ID      uint   `json:"id"`
	FirstName    string `json:"first_name"`
	LastName   string `json:"last_name"`
	DOB   string `json:"dob"`
	Address string `json:"address"`
}
func (b *Student) TableName() string {
	return "student"
}
