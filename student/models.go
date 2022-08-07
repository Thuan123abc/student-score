package student

type Student struct {
	ID        int64
	Name      string
	Address   string
	ScorePrio float64
}
type StudentA struct {
	Student
	MathScore      float64
	PhysicalScore  float64
	ChemistryScore float64
}
type StudentB struct {
	Student
	MathScore      float64
	BioScore       float64
	ChemistryScore float64
}
type StudentC struct {
	Student
	LiteScore    float64
	HistoryScore float64
	GeoScore     float64
}
