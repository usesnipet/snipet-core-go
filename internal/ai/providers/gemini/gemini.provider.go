package gemini

import (
	"context"
	"errors"

	"google.golang.org/genai"

	"github.com/usesnipet/snipet-core-go/internal/ai/model"
)

type Provider struct {
	client *genai.Client
}

func New(apiKey string) (*Provider, error) {
	if apiKey == "" {
		return nil, errors.New("gemini api key is required")
	}

	client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		return nil, err
	}

	return &Provider{client: client}, nil
}

func (p *Provider) Name() string {
	return "gemini"
}

func (p *Provider) TextModel(name string) (model.TextModel, error) {
	if name == "" {
		return nil, errors.New("model name is required")
	}

	return &textModel{
		client:  p.client,
		modelID: name,
	}, nil
}

func (p *Provider) EmbeddingModel(name string) (model.EmbeddingModel, error) {
	if name == "" {
		return nil, errors.New("model name is required")
	}

	return &embeddingModel{
		client:  p.client,
		modelID: name,
	}, nil
}
