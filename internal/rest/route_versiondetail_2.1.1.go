package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/satimoto/go-ocpi/internal/versiondetail"
)

func (rs *RestService) mountVersionDetails() *chi.Mux {
	versionDetailResolver := versiondetail.NewResolver(rs.RepositoryService, rs.ServiceResolver)
	router := chi.NewRouter()

	router.Use(rs.CredentialContextByToken)
	router.Get("/", versionDetailResolver.GetVersionDetail)

	return router
}
