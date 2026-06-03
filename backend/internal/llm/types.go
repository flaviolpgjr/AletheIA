package llm

type PromiseExtraction struct {
	Category   string   `json:"category"`
	Goal       string   `json:"goal"`
	Deadline   string   `json:"deadline"`
	Indicators []string `json:"indicators"`

	Criteria []ExtractedCriterion `json:"criteria"`
}

type ExtractedCriterion struct {
	Key         string `json:"key"`
	Status      string `json:"status"`
	Explanation string `json:"explanation"`
}