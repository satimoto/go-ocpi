package router

import (
	"database/sql"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	router211 "github.com/satimoto/go-ocpi-api/router/v2.1.1"
	"github.com/satimoto/go-ocpi-api/version"
)

func Initialize(d *sql.DB) *chi.Mux {
	repositoryService := db.NewRepositoryService(d)
	router := chi.NewRouter()

	// Set middleware
	router.Use(render.SetContentType(render.ContentTypeJSON), middleware.RedirectSlashes, middleware.Recoverer)

	router.Use(middleware.Timeout(30 * time.Second))

	router.Mount("/", version.New(repositoryService))

	// Adds routes
	router.Route("/2.1.1", func(r chi.Router) {
		r.Mount("/", router211.Routes(repositoryService))
	})

	return router
}
