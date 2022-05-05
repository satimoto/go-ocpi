package cdr

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *CdrResolver) CdrContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if cdrID := chi.URLParam(request, "cdr_id"); cdrID != "" {
			cdr, err := r.Repository.GetCdrByUid(ctx, cdrID)

			if err == nil {
				ctx = context.WithValue(ctx, "cdr", cdr)
				next.ServeHTTP(rw, request.WithContext(ctx))
				return
			}
		}

		render.Render(rw, request, transportation.OCPIErrorUnknownResource(nil))
	})
}
