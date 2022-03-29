package api

import (
	"github.com/go-chi/chi/v5"
)

func (rs *RouterService) mount211() *chi.Mux {
	router := chi.NewRouter()

	//router.Use(credential.CredentialContextByToken(versionResolver.Repository))
	router.Mount("/", rs.mountVersionDetails())
	router.Mount("/credentials", rs.mountCredentials())
	router.Mount("/locations", rs.mountLocations())
	router.Mount("/sessions", rs.mountSessions())
	router.Mount("/tariffs", rs.mountTariffs())
	router.Mount("/tokens", rs.mountTokens())

	return router
}
