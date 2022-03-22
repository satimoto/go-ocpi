package api

import (
	"github.com/go-chi/chi/v5"
	evse "github.com/satimoto/go-ocpi-api/internal/evse/v2.1.1"
)

func (rs *RouterService) mountEvses() *chi.Mux {
	evseResolver := evse.NewResolver(rs.RepositoryService)
	router := chi.NewRouter()

	router.Route("/{evse_id}", func(evseRouter chi.Router) {
		evseRouter.Put("/", evseResolver.UpdateEvse)

		evseContextRouter := evseRouter.With(evseResolver.EvseContext)
		evseContextRouter.Get("/", evseResolver.GetEvse)
		evseContextRouter.Patch("/", evseResolver.UpdateEvse)

		evseContextRouter.Mount("/", rs.mountConnectors())
	})

	return router
}
