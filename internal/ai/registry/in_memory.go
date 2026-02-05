package registry

import (
	"fmt"

	"github.com/usesnipet/snipet-core-go/internal/ai/model"
)

type inMemoryRegistry struct {
	providers map[string]model.Provider
}

func (r *inMemoryRegistry) Register(provider model.Provider) error {
	if provider == nil {
		return fmt.Errorf("provider cannot be nil")
	}

	name := provider.Name()
	if name == "" {
		return fmt.Errorf("provider name cannot be empty")
	}

	if _, exists := r.providers[name]; exists {
		return fmt.Errorf("provider already registered: %s", name)
	}

	r.providers[name] = provider
	return nil
}

func (r *inMemoryRegistry) Provider(name string) (model.Provider, bool) {
	provider, ok := r.providers[name]
	return provider, ok
}

func (r *inMemoryRegistry) List() []model.Provider {
	out := make([]model.Provider, 0, len(r.providers))
	for _, p := range r.providers {
		out = append(out, p)
	}
	return out
}

func NewInMemoryRegistry() model.Registry {
	return &inMemoryRegistry{
		providers: make(map[string]model.Provider),
	}
}
