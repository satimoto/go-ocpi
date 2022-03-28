package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/satimoto/go-ocpi-api/internal/middleware"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1"
)

func (rs *RouterService) mountTokens() *chi.Mux {
	tokenResolver := token.NewResolver(rs.RepositoryService)
	router := chi.NewRouter()

	paginationContextRouter := router.With(middleware.Paginate)
	paginationContextRouter.Get("/", tokenResolver.ListTokens)

	router.Route("/{token_id}", func(tokenRouter chi.Router) {
		tokenContextRouter := tokenRouter.With(tokenResolver.TokenContext)
		tokenContextRouter.Post("/authorize", tokenResolver.AuthorizeToken)
	})

	return router
}
