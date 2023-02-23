package rest

import (
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1"
)

func (s *RestService) mountEvses() *chi.Mux {
	evseResolver := evse.NewResolver(s.RepositoryService, s.ServiceResolver)
	s.evseRestService = evse.NewRestService(evseResolver)

	router := chi.NewRouter()
	router.Use(middleware.Timeout(30 * time.Second))

	router.Route("/{evse_id}", func(evseRouter chi.Router) {
		evseRouter.Put("/", s.evseRestService.UpdateEvse)

		evseContextRouter := evseRouter.With(evseResolver.EvseContext(s.ServiceResolver.SyncService))
		evseContextRouter.Get("/", s.evseRestService.GetEvse)
		evseContextRouter.Patch("/", s.evseRestService.UpdateEvse)

		evseContextRouter.Mount("/", s.mountConnectors())
	})

	return router
}
