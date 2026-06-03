package llm

type Client interface {
	ExtractPromise(text string) (*PromiseExtraction, error)
}