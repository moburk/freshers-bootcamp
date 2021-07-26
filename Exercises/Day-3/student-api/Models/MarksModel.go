package Models

type Marks struct {
	ID uint `json:"id"`
	StudentID uint `json:"student_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Student Student 
	SubjectID uint `json:"subject_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Subject Subject
	Marks int `json:"marks"`
}

func (m *Marks) TableName() string {
	return "marks"
}
