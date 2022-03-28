package token

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
)

func (r *TokenResolver) TokenContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if tokenID := chi.URLParam(request, "token_id"); tokenID != "" {
			token, err := r.Repository.GetTokenByUid(ctx, tokenID)

			if err == nil {
				ctx = context.WithValue(ctx, "token", token)
				next.ServeHTTP(rw, request.WithContext(ctx))
				return
			}
		}

		render.Render(rw, request, ocpi.OCPIErrorUnknownResource(nil))
	})
}
