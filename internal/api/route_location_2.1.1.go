package api

import (
	"github.com/go-chi/chi/v5"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
)

func (rs *RouterService) mountLocations() *chi.Mux {
	locationResolver := location.NewResolver(rs.RepositoryService)
	router := chi.NewRouter()

	router.Route("/{country_code}/{party_id}", func(credentialRouter chi.Router) {
		credentialRouter.Use(locationResolver.CredentialResolver.CredentialContextByPartyAndCountry)

		credentialRouter.Route("/{location_id}", func(locationRouter chi.Router) {
			locationRouter.Put("/", locationResolver.UpdateLocation)

			locationContextRouter := locationRouter.With(locationResolver.LocationContext)
			locationContextRouter.Get("/", locationResolver.GetLocation)
			locationContextRouter.Patch("/", locationResolver.UpdateLocation)

			locationContextRouter.Mount("/", rs.mountEvses())
		})
	})

	return router
}
