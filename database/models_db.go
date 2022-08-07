package database

// Định nghĩa struct StudentModel ánh xạ sang bảng student trong database
type StudentModel struct {
	Mssv           string `gorm:"primarykey"`
	Name           string
	Address        string
	ScorePrio      float64
	Type           string
	MathScore      float64
	PhysicalScore  float64
	BioScore       float64
	ChemistryScore float64
	LiteScore      float64
	HistoryScore   float64
	GeoScore       float64
}

func (StudentModel) TableName() string {
	return "student"
}
