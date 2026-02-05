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
	resp, err := m.client.Models.GenerateContent(ctx, m.modelID, content, nil)
	if err != nil {
		return model.TextResult{}, err
	}

	end := time.Now()

	output := ""
	if resp != nil {
		output = resp.Text()
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

	if resp != nil && resp.UsageMetadata != nil {
		inputT := int(resp.UsageMetadata.PromptTokenCount)
		outputT := int(resp.UsageMetadata.CandidatesTokenCount)
		totalT := int(resp.UsageMetadata.TotalTokenCount)
		result.Usage = &model.Usage{
			InputTokens:  &inputT,
			OutputTokens: &outputT,
			TotalTokens:  &totalT,
		}
	}

	return model.TextResult{
		Result: result,
		Output: output,
		Stream: false,
	}, nil
}

func (m *textModel) Stream(ctx context.Context, req model.TextRequest) (*model.TextStream, error) {
	chunks := make(chan model.StreamChunk)
	done := make(chan model.Result)
	errCh := make(chan error, 1)

	go func() {
		defer close(chunks)
		defer close(done)
		defer close(errCh)

		start := time.Now()
		content := genai.Text(req.Prompt)
		seq := m.client.Models.GenerateContentStream(ctx, m.modelID, content, nil)

		var lastResp *genai.GenerateContentResponse
		for lastResp, err := range seq {
			if err != nil {
				errCh <- err
				return
			}
			if lastResp == nil {
				continue
			}
			t := lastResp.Text()
			if t != "" {
				chunks <- model.StreamChunk{Text: t}
			}
		}

		end := time.Now()
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
		if lastResp != nil && lastResp.UsageMetadata != nil {
			inputT := int(lastResp.UsageMetadata.PromptTokenCount)
			outputT := int(lastResp.UsageMetadata.CandidatesTokenCount)
			totalT := int(lastResp.UsageMetadata.TotalTokenCount)
			result.Usage = &model.Usage{
				InputTokens:  &inputT,
				OutputTokens: &outputT,
				TotalTokens:  &totalT,
			}
		}
		done <- *result
	}()

	return &model.TextStream{
		Chunks: chunks,
		Done:   done,
		Error:  errCh,
	}, nil
}
