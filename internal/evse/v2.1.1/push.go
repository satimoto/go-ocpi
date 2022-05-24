package evse

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *EvseResolver) GetEvse(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	evse := ctx.Value("evse").(db.Evse)
	dto := r.CreateEvseDto(ctx, evse)

	if err := render.Render(rw, request, transportation.OcpiSuccess(dto)); err != nil {
		util.LogOnError("OCPI110", "Error rendering response", err)
		util.LogHttpRequest("OCPI110", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}

func (r *EvseResolver) UpdateEvse(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	location := ctx.Value("location").(db.Location)
	uid := chi.URLParam(request, "evse_id")
	dto := EvseDto{}

	if err := json.NewDecoder(request.Body).Decode(&dto); err != nil {
		util.LogOnError("OCPI111", "Error unmarshalling request", err)
		util.LogHttpRequest("OCPI111", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	evse := r.ReplaceEvse(ctx, location.ID, uid, &dto)

	if evse != nil {
		if dto.Capabilities != nil || dto.Status != nil {
			r.updateLocationAvailability(ctx, evse.ID)
		}

		updateLocationLastUpdatedParams := db.UpdateLocationLastUpdatedParams{
			ID:          location.ID,
			LastUpdated: evse.LastUpdated,
		}
		err := r.Repository.UpdateLocationLastUpdated(ctx, updateLocationLastUpdatedParams)

		if err != nil {
			util.LogOnError("OCPI112", "Error updating evse", err)
			log.Printf("OCPI112: Params=%#v", updateLocationLastUpdatedParams)

			render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
			return
		}
	}

	render.Render(rw, request, transportation.OcpiSuccess(nil))
}
