package gemini

import (
	"context"
	"time"

	"google.golang.org/genai"

	"github.com/usesnipet/snipet-core-go/internal/ai/model"
)

type textModel struct {
	client  *genai.Client
	modelID string
}

func (m *textModel) Generate(ctx context.Context, req model.TextRequest) (model.TextResult, error) {
	start := time.Now()

	content := genai.Text(req.Prompt)

	resp, err := m.client.Models.GenerateContent(
		ctx,
		m.modelID,
		content,
		nil,
	)
	if err != nil {
		return model.TextResult{}, err
	}

	end := time.Now()

	output := ""
	if resp != nil && len(resp.Candidates) > 0 {
		for _, part := range resp.Candidates[0].Content.Parts {
			if t, ok := part.(genai.Text); ok {
				output += string(t)
			}
		}
	}

	result := &model.Result{
		Type: model.ModelTypeText,
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
			InputTokens:  resp.UsageMetadata.PromptTokenCount,
			OutputTokens: resp.UsageMetadata.CandidatesTokenCount,
			TotalTokens:  resp.UsageMetadata.TotalTokenCount,
		}
	}

	return model.TextResult{
		Result: result,
		Output: output,
		Stream: false,
	}, nil
}
