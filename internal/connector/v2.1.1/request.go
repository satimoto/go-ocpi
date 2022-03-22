package connector

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
)

func (r *ConnectorResolver) GetConnector(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	connector := ctx.Value("connector").(db.Connector)
	payload := r.CreateConnectorPayload(ctx, connector)

	if err := render.Render(rw, request, ocpi.OCPISuccess(payload)); err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}
}

func (r *ConnectorResolver) UpdateConnector(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	evse := ctx.Value("evse").(db.Evse)
	uid := chi.URLParam(request, "connector_id")
	payload := ConnectorPayload{}

	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
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
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}

	render.Render(rw, request, ocpi.OCPISuccess(nil))
}
