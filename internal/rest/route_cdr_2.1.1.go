package rest

import (
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	cdr "github.com/satimoto/go-ocpi/internal/cdr/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/version"
)

func (rs *RestService) mountCdrs() *chi.Mux {
	cdrResolver := cdr.NewResolver(rs.RepositoryService, rs.ServiceResolver)
	rs.ServiceResolver.SyncService.AddHandler(version.VERSION_2_1_1, cdrResolver)

	router := chi.NewRouter()
	router.Use(middleware.Timeout(30 * time.Second))
	router.Use(rs.CredentialContextByToken)

	router.Route("/", func(credentialRouter chi.Router) {
		credentialRouter.Post("/", cdrResolver.PostCdr)

		credentialRouter.Route("/{cdr_id}", func(cdrRouter chi.Router) {
			cdrContextRouter := cdrRouter.With(cdrResolver.CdrContext)
			cdrContextRouter.Get("/", cdrResolver.GetCdr)
		})
	})

	return router
}
