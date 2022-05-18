package connector

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *ConnectorResolver) ConnectorContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if connectorID := chi.URLParam(request, "connector_id"); connectorID != "" {
			evse := ctx.Value("evse").(*db.Evse)

			connector, err := r.Repository.GetConnectorByUid(ctx, db.GetConnectorByUidParams{
				EvseID: evse.ID,
				Uid:    connectorID,
			})

			if err == nil {
				ctx = context.WithValue(ctx, "connector", connector)
				next.ServeHTTP(rw, request.WithContext(ctx))
				return
			}
		}

		render.Render(rw, request, transportation.OcpiErrorUnknownResource(nil))
	})
}
