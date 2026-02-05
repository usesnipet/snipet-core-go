package registry

import (
	"fmt"

	"github.com/usesnipet/snipet-core-go/internal/infra/provider"
)

type Registry struct {
	providers map[string]provider.SourceProvider
}

func NewRegistry() *Registry {
	return &Registry{
		providers: make(map[string]provider.SourceProvider),
	}
}

func (r *Registry) Register(p provider.SourceProvider) {
	r.providers[p.Name()] = p
}

func (r *Registry) Get(name string) (provider.SourceProvider, error) {
	p, ok := r.providers[name]
	if !ok {
		return nil, fmt.Errorf("provider %s not found", name)
	}
	return p, nil
}
