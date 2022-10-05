package connector

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *ConnectorResolver) GetConnector(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	connector := ctx.Value("connector").(db.Connector)
	dto := r.CreateConnectorDto(ctx, connector)

	if err := render.Render(rw, request, transportation.OcpiSuccess(dto)); err != nil {
		util.LogOnError("OCPI081", "Error rendering response", err)
		util.LogHttpRequest("OCPI081", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}

func (r *ConnectorResolver) UpdateConnector(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	evse := ctx.Value("evse").(db.Evse)
	uid := chi.URLParam(request, "connector_id")
	connectorDto := dto.ConnectorDto{}

	if err := json.NewDecoder(request.Body).Decode(&connectorDto); err != nil {
		util.LogOnError("OCPI082", "Error unmarshalling request", err)
		util.LogHttpRequest("OCPI082", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	connector := r.ReplaceConnector(ctx, evse, uid, &connectorDto)

	if connector != nil {
		updateEvseLastUpdatedParams := db.UpdateEvseLastUpdatedParams{
			ID:          evse.ID,
			LastUpdated: connector.LastUpdated,
		}
		err := r.Repository.UpdateEvseLastUpdated(ctx, updateEvseLastUpdatedParams)

		if err != nil {
			util.LogOnError("OCPI083", "Error updating evse", err)
			log.Printf("OCPI083: Params=%#v", updateEvseLastUpdatedParams)

			render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
			return
		}

		updateLocationLastUpdatedParams := db.UpdateLocationLastUpdatedParams{
			ID:          evse.LocationID,
			LastUpdated: connector.LastUpdated,
		}
		err = r.Repository.UpdateLocationLastUpdated(ctx, updateLocationLastUpdatedParams)

		if err != nil {
			util.LogOnError("OCPI084", "Error updating location", err)
			log.Printf("OCPI084: Params=%#v", updateEvseLastUpdatedParams)

			render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
			return
		}
	}

	render.Render(rw, request, transportation.OcpiSuccess(nil))
}
