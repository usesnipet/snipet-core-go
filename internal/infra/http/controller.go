package http

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type Controller interface {
	RegisterRoutes(r chi.Router)
}

type ControllersIn struct {
	fx.In

	Controllers []Controller `group:"controllers"`
}
