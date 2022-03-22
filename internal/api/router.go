package api

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
)

type Router interface {
	Handler() *chi.Mux
}

type RouterService struct {
	*db.RepositoryService
}

func NewRouter(d *sql.DB) Router {
	return &RouterService{
		RepositoryService: db.NewRepositoryService(d),
	}
}

func (rs *RouterService) Handler() *chi.Mux {
	router := chi.NewRouter()

	// Set middleware
	router.Use(render.SetContentType(render.ContentTypeJSON), middleware.RedirectSlashes, middleware.Recoverer)
	router.Use(middleware.Timeout(30 * time.Second))

	router.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Add routes
	router.Mount("/", rs.mountVersions())
	router.Mount("/2.1.1", rs.mount211())

	fmt.Println(docgen.MarkdownRoutesDoc(router, docgen.MarkdownOpts{
		ProjectPath: "github.com/go-chi/chi",
		Intro:       "Welcome to the chi/_examples/rest generated docs.",
	}))

	return router
}
