package domain

type CriterionStatus string

const (
	CriterionStatusYes     CriterionStatus = "yes"
	CriterionStatusPartial CriterionStatus = "partial"
	CriterionStatusNo      CriterionStatus = "no"
)

type Criterion struct {
	Key         string
	Name        string
	Weight      int
	Status      CriterionStatus
	Score       float64
	Explanation string
}

type PublicSource struct {
	Name        string
	Description string
}

type Analysis struct {
	Summary    string
	Score      int
	Confidence int
	Criteria   []Criterion
	Risks      []string
	Sources    []PublicSource
}