package api

import (
	"github.com/go-chi/chi/v5"
	cdr "github.com/satimoto/go-ocpi-api/internal/cdr/v2.1.1"
)

func (rs *RouterService) mountCdrs() *chi.Mux {
	cdrResolver := cdr.NewResolver(rs.RepositoryService)
	router := chi.NewRouter()

	router.Route("/", func(credentialRouter chi.Router) {
		credentialRouter.Post("/", cdrResolver.PostCdr)

		credentialRouter.Route("/{cdr_id}", func(cdrRouter chi.Router) {
			cdrContextRouter := cdrRouter.With(cdrResolver.CdrContext)
			cdrContextRouter.Get("/", cdrResolver.GetCdr)
		})
	})

	return router
}
