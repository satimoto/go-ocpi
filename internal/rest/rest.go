package rest

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/edjumacator/chi-prometheus"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

type Rest interface {
	StartRest(context.Context, *sync.WaitGroup)
}

type RestService struct {
	*db.RepositoryService
	*http.Server
}

func NewRest(d *sql.DB) Rest {
	restService := &RestService{
		RepositoryService: db.NewRepositoryService(d),
	}
	restService.Server = &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("REST_PORT")),
		Handler: restService.handler(),
	}

	return restService
}

func (rs *RestService) StartRest(ctx context.Context, waitGroup *sync.WaitGroup) {
	log.Printf("Starting Rest service")
	waitGroup.Add(1)

	go rs.listenAndServe()

	go func() {
		<-ctx.Done()
		log.Printf("Shutting down Rest service")

		rs.shutdown()

		log.Printf("Rest service shut down")
		waitGroup.Done()
	}()
}

func (rs *RestService) listenAndServe() {
	err := rs.Server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Printf("Error in Rest service: %v", err)
	}
}

func (rs *RestService) shutdown() {
	timeout := util.ParseInt32(os.Getenv("SHUTDOWN_TIMEOUT"), 20)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	err := rs.Server.Shutdown(ctx)

	if err != nil {
		log.Printf("Error shutting down Rest service: %v", err)
	}
}

func (rs *RestService) handler() *chi.Mux {
	router := chi.NewRouter()

	// Set middleware
	router.Use(render.SetContentType(render.ContentTypeJSON), middleware.RedirectSlashes, middleware.Recoverer)
	router.Use(middleware.Timeout(30 * time.Second))
	router.Use(chiprometheus.NewMiddleware("ocpi"))

	router.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Add routes
	router.Mount("/", rs.mountVersions())
	router.Mount("/metrics", promhttp.Handler())
	router.Mount("/2.1.1", rs.mount211())

	return router
}
