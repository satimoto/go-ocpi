package rest

import (
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	ocpiMiddleware "github.com/satimoto/go-ocpi/internal/middleware"
	token "github.com/satimoto/go-ocpi/internal/token/v2.1.1"
)

func (rs *RestService) mountTokens() *chi.Mux {
	tokenResolver := token.NewResolver(rs.RepositoryService, rs.ServiceResolver)
	router := chi.NewRouter()
	router.Use(rs.CredentialContextByToken)

	paginationContextRouter := router.With(ocpiMiddleware.Paginate)
	paginationContextRouter.Use(middleware.Timeout(30 * time.Second))
	paginationContextRouter.Get("/", tokenResolver.ListTokens)

	router.Route("/{token_id}", func(tokenRouter chi.Router) {
		tokenContextRouter := tokenRouter.With(tokenResolver.TokenContext)
		tokenContextRouter.Post("/authorize", tokenResolver.AuthorizeToken)
	})

	return router
}
