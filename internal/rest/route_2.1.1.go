package rest

import (
	"github.com/go-chi/chi/v5"
)

func (rs *RestService) mount211() *chi.Mux {
	router := chi.NewRouter()

	router.Mount("/", rs.mountVersionDetails())
	router.Mount("/cdrs", rs.mountCdrs())
	router.Mount("/commands", rs.mountCommands())
	router.Mount("/credentials", rs.mountCredentials())
	router.Mount("/locations", rs.mountLocations())
	router.Mount("/sessions", rs.mountSessions())
	router.Mount("/tariffs", rs.mountTariffs())
	router.Mount("/tokens", rs.mountTokens())

	return router
}
