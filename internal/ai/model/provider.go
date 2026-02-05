package model

type Provider interface {
	Name() string

	TextModel(name string) (TextModel, error)
	EmbeddingModel(name string) (EmbeddingModel, error)
}
