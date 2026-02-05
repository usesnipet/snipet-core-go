package model

import "context"

// Embedding model
type EmbeddingModel interface {
	Embed(ctx context.Context, req EmbeddingRequest) (EmbeddingResult, error)
	EmbedBatch(ctx context.Context, req BatchEmbeddingRequest) (BatchEmbeddingResult, error)
}

// Text model
type TextModel interface {
	Generate(ctx context.Context, req TextRequest) (TextResult, error)
	Stream(ctx context.Context, req TextRequest) (*TextStream, error)
}
