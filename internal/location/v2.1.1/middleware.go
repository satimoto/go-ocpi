package location

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *LocationResolver) LocationContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if locationID := chi.URLParam(request, "location_id"); locationID != "" {
			location, err := r.Repository.GetLocationByUid(ctx, locationID)

			if err == nil {
				ctx = context.WithValue(ctx, "location", location)
				next.ServeHTTP(rw, request.WithContext(ctx))
				return
			}
		}

		render.Render(rw, request, transportation.OCPIErrorUnknownResource(nil))
	})
}