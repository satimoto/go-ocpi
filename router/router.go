package router

import (
	"database/sql"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	router211 "github.com/satimoto/go-ocpi-api/router/v2.1.1"
)

func Initialize(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()

	// Set middleware
	router.Use(render.SetContentType(render.ContentTypeJSON), middleware.RedirectSlashes, middleware.Recoverer)

	router.Use(middleware.Timeout(30 * time.Second))

	// Adds routes
	router.Route("/2.1.1", func(r chi.Router) {
		r.Mount("/", router211.Routes(db))
	})

	return router
}
