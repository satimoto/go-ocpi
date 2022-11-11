package rest

import (
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	coreLocation "github.com/satimoto/go-ocpi/internal/location"
	location "github.com/satimoto/go-ocpi/internal/location/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/version"
)

func (rs *RestService) mountLocations() *chi.Mux {
	locationResolver := location.NewResolver(rs.RepositoryService, rs.ServiceResolver)
	rs.ServiceResolver.SyncService.AddHandler(version.VERSION_2_1_1, coreLocation.IDENTIFIER, locationResolver)

	router := chi.NewRouter()
	router.Use(middleware.Timeout(30 * time.Second))
	router.Use(rs.CredentialContextByToken)

	router.Route("/{country_code}/{party_id}/{location_id}", func(locationRouter chi.Router) {
		locationRouter.Put("/", locationResolver.UpdateLocation)

		locationContextRouter := locationRouter.With(locationResolver.LocationContext(rs.ServiceResolver.SyncService))
		locationContextRouter.Get("/", locationResolver.GetLocation)
		locationContextRouter.Patch("/", locationResolver.UpdateLocation)

		locationContextRouter.Mount("/", rs.mountEvses())
	})

	return router
}
