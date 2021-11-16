package connector

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/rest"
)

func (r *ConnectorResolver) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/{connectorId}", func(connectorRouter chi.Router) {
		connectorRouter.Put("/", r.updateConnector)

		connectorContextRouter := connectorRouter.With(r.connectorContext)
		connectorContextRouter.Get("/", r.getConnector)
		connectorContextRouter.Patch("/", r.updateConnector)
	})

	return router
}

func (r *ConnectorResolver) connectorContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if connectorID := chi.URLParam(request, "connectorID"); connectorID != "" {
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

		render.Render(rw, request, rest.OCPIErrorUnknownResource(nil))
	})
}

func (r *ConnectorResolver) getConnector(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	connector := ctx.Value("connector").(db.Connector)
	payload := r.CreateConnectorPayload(ctx, connector)

	if err := render.Render(rw, request, rest.OCPISuccess(payload)); err != nil {
		render.Render(rw, request, rest.OCPIServerError(nil, err.Error()))
	}
}

func (r *ConnectorResolver) updateConnector(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	evse := ctx.Value("evse").(db.Evse)
	uid := chi.URLParam(request, "connectorID")
	payload := ConnectorPayload{}

	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		render.Render(rw, request, rest.OCPIServerError(nil, err.Error()))
	}

	connector := r.ReplaceConnector(ctx, evse.ID, uid, &payload)

	err := r.Repository.UpdateEvseLastUpdated(ctx, db.UpdateEvseLastUpdatedParams{
		ID:          evse.ID,
		LastUpdated: connector.LastUpdated,
	})

	err = r.Repository.UpdateLocationLastUpdated(ctx, db.UpdateLocationLastUpdatedParams{
		ID:          evse.LocationID,
		LastUpdated: connector.LastUpdated,
	})

	if err != nil {
		render.Render(rw, request, rest.OCPIServerError(nil, err.Error()))
	}

	render.Render(rw, request, rest.OCPISuccess(nil))
}
