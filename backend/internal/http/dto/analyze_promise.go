package dto

type AnalyzePromiseRequest struct {
	Text string `json:"text"`
}

type AnalyzePromiseResponse struct {
	Summary string   `json:"summary"`
	Score   int      `json:"score"`
	Risks   []string `json:"risks"`
}
