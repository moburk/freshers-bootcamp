package Models

type Subject struct {
	ID uint `json:"id"`
	SubjectName string `json:"subject_name"`
}

func (s *Subject) TableName() string {
	return "subject"
}
