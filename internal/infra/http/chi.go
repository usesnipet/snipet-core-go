package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/usesnipet/snipet-core-go/internal/config"
	"go.uber.org/fx"
)

func NewChi() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	return r
}

var Module = fx.Module(
	"http",
	fx.Provide(NewChi),
	fx.Invoke(func(
		lc fx.Lifecycle,
		r *chi.Mux,
		controllersIn ControllersIn,
	) {
		r.Route("/api", func(r chi.Router) {
			for _, controller := range controllersIn.Controllers {
				controller.RegisterRoutes(r)
			}
		})
		server := &http.Server{
			Addr:    config.GetEnv().APP_PORT,
			Handler: r,
		}

		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go server.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return server.Shutdown(ctx)
			},
		})
	}),
)
