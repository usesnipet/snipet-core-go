package model

type ModelType string

const (
	ModelTypeText      ModelType = "text"
	ModelTypeEmbedding ModelType = "embedding"
)

type ModelInfo struct {
	Name     string // ex: gpt-4o-mini
	Provider string // ex: openai, anthropic, ollama
}

type Timing struct {
	StartedAt  int64 // unix millis
	FinishedAt int64 // unix millis
	DurationMs int64 // FinishedAt - StartedAt
}
type Usage struct {
	InputTokens  *int
	OutputTokens *int
	TotalTokens  *int
}

type Result struct {
	Type ModelType

	Model  ModelInfo
	Timing Timing
	Usage  *Usage
}

type TextResult struct {
	*Result
	Output string

	Stream bool
}

type EmbeddingResult struct {
	*Result
	Output []float64
}
type BatchEmbeddingResult struct {
	*Result
	Output [][]float64
}
