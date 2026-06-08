package llm

type geminiRequest struct {
	Contents []geminiContent `json:"contents"`
}

type geminiResponse struct {
	Candidates []geminiCandidate `json:"candidates"`
}

type geminiCandidate struct {
	Content geminiContent `json:"content"`
}

type geminiContent struct {
	Parts []geminiPart `json:"parts"`
	Role  string       `json:"role,omitempty"`
}

type geminiPart struct {
	Text string `json:"text"`
}