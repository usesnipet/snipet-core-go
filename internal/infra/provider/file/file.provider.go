package file

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/usesnipet/snipet-core-go/internal/entity"
	"github.com/usesnipet/snipet-core-go/internal/infra/provider"
	"github.com/usesnipet/snipet-core-go/internal/infra/storage"
)

type Provider struct {
	storage *storage.Service
}

func New(storage *storage.Service) *Provider {
	return &Provider{storage: storage}
}

func (p *Provider) Name() string {
	return "file"
}

// ValidateConfig checks that config only contains allowed keys and valid types.
// Allowed: "prefix" (optional, must be string if present).
func (p *Provider) ValidateConfig(config map[string]any) error {
	return nil
}

// TestConnection verifies access to storage by listing objects (with optional prefix).
func (p *Provider) TestConnection(ctx context.Context, ks entity.KnowledgeSource, config map[string]any) error {
	_, err := p.storage.List(ctx, &ks.ID)
	return err
}

// List get a slice of object ref in the file storage
func (p *Provider) List(
	ctx context.Context,
	ks entity.KnowledgeSource,
	raw map[string]any,
) ([]provider.ObjectRef, error) {
	err := p.ValidateConfig(raw)
	if err != nil {
		return nil, err
	}
	objects, err := p.storage.List(ctx, &ks.ID)
	if err != nil {
		return nil, err
	}

	refs := make([]provider.ObjectRef, 0, len(objects))

	for _, obj := range objects {
		refs = append(refs, mapObject(
			p.Name(),
			ks.ID,
			aws.ToString(&obj.Path),
			obj.Size,
			*obj.LastModified,
		))
	}

	return refs, nil
}

// Load get the object in the file storage
func (p *Provider) Read(
	ctx context.Context,
	ref provider.ObjectRef,
	raw map[string]any,
) (*provider.Object, error) {
	err := p.ValidateConfig(raw)
	if err != nil {
		return nil, err
	}
	reader, file, err := p.storage.Get(ctx, ref.Path)
	if err != nil {
		return nil, err
	}

	return &provider.Object{
		ObjectRefID: ref.ID,
		MimeType:    *file.ContentType,
		Size:        *file.Size,
		Stream:      reader,
	}, nil
}
