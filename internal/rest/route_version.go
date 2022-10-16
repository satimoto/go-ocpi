package rest

import (
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/satimoto/go-ocpi/internal/version"
)

func (rs *RestService) mountVersions() *chi.Mux {
	versionResolver := version.NewResolver(rs.RepositoryService, rs.ServiceResolver.OcpiService)
	
	router := chi.NewRouter()
	router.Use(middleware.Timeout(30 * time.Second))
	router.Use(middleware.Logger)
	router.Use(rs.CredentialContextByToken)
	router.Get("/", versionResolver.GetVersions)

	return router
}
