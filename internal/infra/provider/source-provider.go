package provider

import (
	"context"

	"github.com/usesnipet/snipet-core-go/internal/entity"
)

type SourceProvider interface {
	Name() string

	ValidateConfig(config map[string]any) error
	TestConnection(ctx context.Context, ks entity.KnowledgeSource, config map[string]any) error

	List(
		ctx context.Context,
		source entity.KnowledgeSource,
		config map[string]any,
	) ([]ObjectRef, error)

	Read(
		ctx context.Context,
		refID ObjectRef,
		config map[string]any,
	) (*Object, error)
}
