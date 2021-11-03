package router

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/satimoto/go-datastore/db"
	location "github.com/satimoto/go-ocpi-api/location/v2.1.1"
)

// Routes initializes the handlers for the router
func Routes(d *sql.DB) *chi.Mux {
	repositoryService := db.NewRepositoryService(d)
	router := chi.NewRouter()
	router.Mount("/locations", location.New(repositoryService))

	return router
}
