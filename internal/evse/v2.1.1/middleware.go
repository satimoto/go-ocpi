package evse

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

func (r *EvseResolver) EvseContext(syncService *sync.SyncService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
			requestCtx := request.Context()

			if evseID := chi.URLParam(request, "evse_id"); evseID != "" {
				evse, err := r.Repository.GetEvseByUid(requestCtx, evseID)

				if err == nil {
					requestCtx = context.WithValue(requestCtx, "evse", evse)
					next.ServeHTTP(rw, request.WithContext(requestCtx))
					return
				}
			}

			if request.Method == http.MethodPatch {
				countryCode := util.NilString(chi.URLParam(request, "country_code"))
				partyID := util.NilString(chi.URLParam(request, "party_id"))
				credential := middleware.GetCredential(requestCtx)

				go syncService.SynchronizeCredential(*credential, true, true, nil, countryCode, partyID)
			}

			render.Render(rw, request, transportation.OcpiErrorUnknownResource(nil))
		})
	}
}
