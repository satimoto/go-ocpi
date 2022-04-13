package rest

import (
	"github.com/go-chi/chi/v5"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail/v2.1.1"
)

func (rs *RestService) mountVersionDetails() *chi.Mux {
	versionDetailResolver := versiondetail.NewResolver(rs.RepositoryService)
	router := chi.NewRouter()

	router.Use(rs.CredentialContextByToken)
	router.Get("/", versionDetailResolver.GetVersionDetail)

	return router
}
