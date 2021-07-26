package Models

type Marks struct {
	//gorm.Model
	ID uint `json:"id"`
	StudentID uint `json:"student_id"` //sql:"REFERENCES students(id) ON DELETE CASCADE ON UPDATE CASCADE"`
	Student Student `gorm:"foreignkey:student_id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SubjectID uint `json:"subject_id"`// sql:"REFERENCES subjects(id) ON DELETE CASCADE ON UPDATE CASCADE"`
	Subject Subject `gorm:"foreignkey:subject_id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Marks int `json:"marks"`
}

func (m *Marks) TableName() string {
	return "marks"
}
