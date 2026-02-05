package http_server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

type Controller interface {
	RegisterRoutes(r chi.Router)
}

type ControllersIn struct {
	fx.In

	Controllers []Controller `group:"controllers"`
}

type BaseController struct {
	validate *validator.Validate
}

func NewBaseController(validate *validator.Validate) *BaseController {
	return &BaseController{validate: validate}
}
