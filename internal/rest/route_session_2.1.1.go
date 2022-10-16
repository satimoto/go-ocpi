package rest

import (
	"github.com/go-chi/chi/v5"
	session "github.com/satimoto/go-ocpi/internal/session/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/version"
)

func (rs *RestService) mountSessions() *chi.Mux {
	sessionResolver := session.NewResolver(rs.RepositoryService, rs.ServiceResolver)
	rs.ServiceResolver.SyncService.AddHandler(version.VERSION_2_1_1, sessionResolver)

	router := chi.NewRouter()
	router.Use(rs.CredentialContextByToken)

	router.Route("/{country_code}/{party_id}/{session_id}", func(sessionRouter chi.Router) {
		sessionRouter.Put("/", sessionResolver.UpdateSession)

		sessionContextRouter := sessionRouter.With(sessionResolver.SessionContext)
		sessionContextRouter.Get("/", sessionResolver.GetSession)
		sessionContextRouter.Patch("/", sessionResolver.UpdateSession)
	})

	return router
}
