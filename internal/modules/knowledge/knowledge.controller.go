package knowledge

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service}
}

func (c *Controller) RegisterRoutes(r chi.Router) {
	r.Route("/knowledge", func(r chi.Router) {
		r.Get("/", c.list)
	})
}

func (c *Controller) list(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}
