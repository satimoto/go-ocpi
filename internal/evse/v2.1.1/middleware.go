package evse

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *EvseResolver) EvseContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if evseID := chi.URLParam(request, "evse_id"); evseID != "" {
			evse, err := r.Repository.GetEvseByUid(ctx, evseID)

			if err == nil {
				ctx = context.WithValue(ctx, "evse", evse)
				next.ServeHTTP(rw, request.WithContext(ctx))
				return
			}
		}

		render.Render(rw, request, transportation.OcpiErrorUnknownResource(nil))
	})
}
