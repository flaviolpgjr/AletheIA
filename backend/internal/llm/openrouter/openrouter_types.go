package openrouter

type openRouterRequest struct {
	Model          string                   `json:"model"`
	Messages       []openRouterMessage      `json:"messages"`
	Temperature    float64                  `json:"temperature"`
	ResponseFormat openRouterResponseFormat `json:"response_format"`
}

type openRouterMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openRouterResponseFormat struct {
	Type string `json:"type"`
}

type openRouterResponse struct {
	Choices []openRouterChoice `json:"choices"`
}

type openRouterChoice struct {
	Message openRouterMessage `json:"message"`
}