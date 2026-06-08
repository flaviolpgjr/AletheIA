package llm

type PromiseExtraction struct {
	Summary    string   `json:"summary"`
	Category   string   `json:"category"`
	Goal       string   `json:"goal"`
	Deadline   string   `json:"deadline"`
	Indicators []string `json:"indicators"`
	Risks      []string `json:"risks"`

	Criteria []ExtractedCriterion `json:"criteria"`
}

type ExtractedCriterion struct {
	Key         string `json:"key"`
	Status      string `json:"status"`
	Explanation string `json:"explanation"`
}