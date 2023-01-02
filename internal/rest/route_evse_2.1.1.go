package rest

import (
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1"
)

func (rs *RestService) mountEvses() *chi.Mux {
	evseResolver := evse.NewResolver(rs.RepositoryService, rs.ServiceResolver)

	router := chi.NewRouter()
	router.Use(middleware.Timeout(30 * time.Second))

	router.Route("/{evse_id}", func(evseRouter chi.Router) {
		evseRouter.Put("/", evseResolver.UpdateEvse)

		evseContextRouter := evseRouter.With(evseResolver.EvseContext(rs.ServiceResolver.SyncService))
		evseContextRouter.Get("/", evseResolver.GetEvse)
		evseContextRouter.Patch("/", evseResolver.UpdateEvse)

		evseContextRouter.Mount("/", rs.mountConnectors())
	})

	return router
}
