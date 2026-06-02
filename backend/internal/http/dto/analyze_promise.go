package dto

type AnalyzePromiseRequest struct {
	Text string `json:"text"`
}

type CriterionResponse struct {
	Key         string  `json:"key"`
	Name        string  `json:"name"`
	Weight      int     `json:"weight"`
	Status      string  `json:"status"`
	Score       float64 `json:"score"`
	Explanation string  `json:"explanation"`
}

type AnalyzePromiseResponse struct {
	Summary    string              `json:"summary"`
	Score      int                 `json:"score"`
	Confidence int                 `json:"confidence"`
	Criteria   []CriterionResponse `json:"criteria"`
	Risks      []string            `json:"risks"`
}