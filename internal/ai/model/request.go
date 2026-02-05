package model

// Embedding request
type EmbeddingRequest struct {
	Input string
}
type BatchEmbeddingRequest struct {
	Inputs []string
}

// Text request
type TextRequest struct {
	Prompt   string
	System   string
	Messages []Message

	Stream bool
}
type Message struct {
	Role    string // system | user | assistant
	Content string
}
