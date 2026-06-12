package llm

type PublicSource struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PromiseExtraction struct {
	Summary    string   `json:"summary"`
	Category   string   `json:"category"`
	Goal       string   `json:"goal"`
	Deadline   string   `json:"deadline"`

	TargetValue float64 `json:"target_value"`
	TargetUnit  string  `json:"target_unit"`

	Indicators []string `json:"indicators"`
	Risks      []string `json:"risks"`

	SuggestedSources []PublicSource `json:"suggested_sources"`

	Criteria []ExtractedCriterion `json:"criteria"`
}

type ExtractedCriterion struct {
	Key         string `json:"key"`
	Status      string `json:"status"`
	Explanation string `json:"explanation"`
}