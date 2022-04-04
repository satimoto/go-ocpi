package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/satimoto/go-ocpi-api/internal/version"
)

func (rs *RouterService) mountVersions() *chi.Mux {
	versionResolver := version.NewResolver(rs.RepositoryService)
	router := chi.NewRouter()
	
	router.Use(rs.CredentialContextByToken)
	router.Get("/", versionResolver.GetVersions)

	return router
}