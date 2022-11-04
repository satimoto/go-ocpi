package evse

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *EvseResolver) GetEvse(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	evse := ctx.Value("evse").(db.Evse)
	dto := r.CreateEvseDto(ctx, evse)

	if err := render.Render(rw, request, transportation.OcpiSuccess(dto)); err != nil {
		metrics.RecordError("OCPI110", "Error rendering response", err)
		util.LogHttpRequest("OCPI110", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}

func (r *EvseResolver) UpdateEvse(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	location := ctx.Value("location").(db.Location)
	uid := chi.URLParam(request, "evse_id")
	evseDto := dto.EvseDto{}

	if err := json.NewDecoder(request.Body).Decode(&evseDto); err != nil {
		metrics.RecordError("OCPI111", "Error unmarshalling request", err)
		util.LogHttpRequest("OCPI111", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	evse := r.ReplaceEvse(ctx, location.ID, uid, &evseDto)

	if evse != nil {
		if evseDto.Capabilities != nil || evseDto.Status != nil {
			r.updateLocationAvailability(ctx, evse.ID)
		}

		updateLocationLastUpdatedParams := db.UpdateLocationLastUpdatedParams{
			ID:          location.ID,
			LastUpdated: evse.LastUpdated,
		}
		err := r.Repository.UpdateLocationLastUpdated(ctx, updateLocationLastUpdatedParams)

		if err != nil {
			metrics.RecordError("OCPI112", "Error updating evse", err)
			log.Printf("OCPI112: Params=%#v", updateLocationLastUpdatedParams)

			render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
			return
		}
	}

	render.Render(rw, request, transportation.OcpiSuccess(nil))
}
