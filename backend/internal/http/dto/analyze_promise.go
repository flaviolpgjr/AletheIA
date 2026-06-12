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

type PublicSourceResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EvidenceResponse struct {
	Source      string  `json:"source"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	URL         string  `json:"url"`

	Indicator   string  `json:"indicator"`
	Value       float64 `json:"value"`
	Unit        string  `json:"unit"`
	Reference   string  `json:"reference"`
}

type AnalyzePromiseResponse struct {
	Summary     string                 `json:"summary"`
	Score       int                    `json:"score"`
	Confidence  int                    `json:"confidence"`

	TargetValue float64                `json:"target_value"`
	TargetUnit  string                 `json:"target_unit"`

	Criteria []CriterionResponse       `json:"criteria"`
	Risks    []string                  `json:"risks"`
	Sources  []PublicSourceResponse    `json:"sources"`
	Evidence []EvidenceResponse        `json:"evidence"`
}