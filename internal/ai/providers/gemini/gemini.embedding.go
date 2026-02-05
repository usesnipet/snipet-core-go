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

	resp, err := m.client.Models.EmbedContent(
		ctx,
		m.modelID,
		genai.Text(req.Input),
		nil,
	)
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

	if resp.UsageMetadata != nil {
		result.Usage = &model.Usage{
			InputTokens: resp.UsageMetadata.PromptTokenCount,
			TotalTokens: resp.UsageMetadata.TotalTokenCount,
		}
	}

	return model.EmbeddingResult{
		Result: result,
		Output: resp.Embedding.Values,
	}, nil
}
