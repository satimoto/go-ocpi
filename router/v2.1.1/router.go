package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/satimoto/go-datastore/db"
	location "github.com/satimoto/go-ocpi-api/location/v2.1.1"
	version "github.com/satimoto/go-ocpi-api/version/v2.1.1"
)

// Routes initializes the handlers for the router
func Routes(repositoryService *db.RepositoryService) *chi.Mux {
	router := chi.NewRouter()
	router.Mount("/", version.New(repositoryService))
	router.Mount("/locations", location.New(repositoryService))

	return router
}
