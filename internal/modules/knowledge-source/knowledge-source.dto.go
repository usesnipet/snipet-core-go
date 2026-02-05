package knowledge_source

import "github.com/usesnipet/snipet-core-go/internal/entity"

// CreateDTO is the request body for creating a knowledge source.
type CreateDTO struct {
	Name         string              `json:"name" validate:"required,max=255"`
	Provider     string              `json:"provider" validate:"required,max=50"`
	ProviderType entity.ProviderType `json:"provider_type" validate:"required,oneof=internal external"`
	Config       map[string]any      `json:"config" validate:"required"`
	UseRAG       *bool               `json:"use_rag"`
	RAGStrategy  entity.RAGStrategy  `json:"rag_strategy" validate:"omitempty,oneof=webhook cron manual"`
	RAGConfig    map[string]any      `json:"rag_config"`
	Status       entity.SourceStatus `json:"status" validate:"omitempty,oneof=active paused error"`
}

// UpdateDTO is the request body for updating a knowledge source (all fields optional).
type UpdateDTO struct {
	Name         *string              `json:"name" validate:"omitempty,max=255"`
	Provider     *string              `json:"provider" validate:"omitempty,max=50"`
	ProviderType *entity.ProviderType `json:"provider_type" validate:"omitempty,oneof=internal external"`
	Config       map[string]any       `json:"config"`
	UseRAG       *bool                `json:"use_rag"`
	RAGStrategy  *entity.RAGStrategy  `json:"rag_strategy" validate:"omitempty,oneof=webhook cron manual"`
	RAGConfig    map[string]any       `json:"rag_config"`
	Status       *entity.SourceStatus `json:"status" validate:"omitempty,oneof=active paused error"`
}
