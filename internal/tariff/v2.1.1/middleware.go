package tariff

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/middleware"
	"github.com/satimoto/go-ocpi/internal/sync"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *TariffResolver) TariffContext(syncService *sync.SyncService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
			requestCtx := request.Context()

			if tariffID := chi.URLParam(request, "tariff_id"); tariffID != "" {
				tariff, err := r.Repository.GetTariffByUid(requestCtx, tariffID)

				if err == nil {
					requestCtx = context.WithValue(requestCtx, "tariff", tariff)
					next.ServeHTTP(rw, request.WithContext(requestCtx))
					return
				}
			}

			if request.Method == http.MethodPatch {
				countryCode := util.NilString(chi.URLParam(request, "country_code"))
				partyID := util.NilString(chi.URLParam(request, "party_id"))
				credential := middleware.GetCredential(requestCtx)

				go syncService.SynchronizeCredential(*credential, nil, countryCode, partyID)
			}

			render.Render(rw, request, transportation.OcpiErrorUnknownResource(nil))
		})
	}
}
