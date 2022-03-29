package api

import (
	"github.com/go-chi/chi/v5"
	session "github.com/satimoto/go-ocpi-api/internal/session/v2.1.1"
)

func (rs *RouterService) mountSessions() *chi.Mux {
	sessionResolver := session.NewResolver(rs.RepositoryService)
	router := chi.NewRouter()

	router.Route("/{country_code}/{party_id}", func(credentialRouter chi.Router) {
		credentialRouter.Use(sessionResolver.CredentialResolver.CredentialContextByPartyAndCountry)

		credentialRouter.Route("/{session_id}", func(sessionRouter chi.Router) {
			sessionRouter.Put("/", sessionResolver.UpdateSession)

			sessionContextRouter := sessionRouter.With(sessionResolver.SessionContext)
			sessionContextRouter.Get("/", sessionResolver.GetSession)
			sessionContextRouter.Patch("/", sessionResolver.UpdateSession)
		})
	})

	return router
}
