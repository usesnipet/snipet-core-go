package model

import "context"

// Embedding model
type EmbeddingModel interface {
	Embed(ctx context.Context, req EmbeddingRequest) (Result, error)
	EmbedBatch(ctx context.Context, req BatchEmbeddingRequest) (Result, error)
}

// Text model
type TextModel interface {
	Generate(ctx context.Context, req TextRequest) (Result, error)
	Stream(ctx context.Context, req TextRequest) (*TextStream, error)
}
