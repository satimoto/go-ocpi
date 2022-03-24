package tariff

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
)

func (r *TariffResolver) TariffContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if tariffID := chi.URLParam(request, "tariff_id"); tariffID != "" {
			tariff, err := r.Repository.GetTariffByUid(ctx, tariffID)

			if err == nil {
				ctx = context.WithValue(ctx, "tariff", tariff)
				next.ServeHTTP(rw, request.WithContext(ctx))
				return
			}
		}

		render.Render(rw, request, ocpi.OCPIErrorUnknownResource(nil))
	})
}