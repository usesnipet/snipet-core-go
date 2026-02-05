package gemini

import (
	"context"
	"time"

	"google.golang.org/genai"

	"github.com/usesnipet/snipet-core-go/internal/ai/model"
)

type embeddingModel struct {
	client  *genai.Client
	modelID string
}

func (m *embeddingModel) Embed(ctx context.Context, req model.EmbeddingRequest) (model.EmbeddingResult, error) {
	start := time.Now()

	contents := genai.Text(req.Input)
	resp, err := m.client.Models.EmbedContent(ctx, m.modelID, contents, nil)
	if err != nil {
		return model.EmbeddingResult{}, err
	}

	end := time.Now()

	result := &model.Result{
		Type: model.ModelTypeEmbedding,
		Model: model.ModelInfo{
			Name:     m.modelID,
			Provider: "gemini",
		},
		Timing: model.Timing{
			StartedAt:  start.UnixMilli(),
			FinishedAt: end.UnixMilli(),
			DurationMs: end.Sub(start).Milliseconds(),
		},
	}

	// EmbedContentResponse has Embeddings[]; single request returns one embedding
	var output []float64
	if len(resp.Embeddings) > 0 && resp.Embeddings[0] != nil {
		for _, v := range resp.Embeddings[0].Values {
			output = append(output, float64(v))
		}
	}

	return model.EmbeddingResult{
		Result: result,
		Output: output,
	}, nil
}

func (m *embeddingModel) EmbedBatch(ctx context.Context, req model.BatchEmbeddingRequest) (model.BatchEmbeddingResult, error) {
	start := time.Now()

	var contents []*genai.Content
	for _, input := range req.Inputs {
		contents = append(contents, genai.Text(input)...)
	}
	resp, err := m.client.Models.EmbedContent(ctx, m.modelID, contents, nil)
	if err != nil {
		return model.BatchEmbeddingResult{}, err
	}

	end := time.Now()

	result := &model.Result{
		Type: model.ModelTypeEmbedding,
		Model: model.ModelInfo{
			Name:     m.modelID,
			Provider: "gemini",
		},
		Timing: model.Timing{
			StartedAt:  start.UnixMilli(),
			FinishedAt: end.UnixMilli(),
			DurationMs: end.Sub(start).Milliseconds(),
		},
	}

	output := make([][]float64, 0, len(resp.Embeddings))
	for _, emb := range resp.Embeddings {
		if emb == nil {
			output = append(output, nil)
			continue
		}
		vec := make([]float64, 0, len(emb.Values))
		for _, v := range emb.Values {
			vec = append(vec, float64(v))
		}
		output = append(output, vec)
	}

	return model.BatchEmbeddingResult{
		Result: result,
		Output: output,
	}, nil
}
