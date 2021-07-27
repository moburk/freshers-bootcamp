package Models

type Subject struct {
	ID uint `json:"id"`
	//gorm.Model
	SubjectName string `json:"subject_name"`
}

func (s *Subject) TableName() string {
	return "subjects"
}
