package evse

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *EvseResolver) GetEvse(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	evse := ctx.Value("evse").(db.Evse)
	dto := r.CreateEvseDto(ctx, evse)

	if err := render.Render(rw, request, transportation.OcpiSuccess(dto)); err != nil {
		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}

func (r *EvseResolver) UpdateEvse(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	location := ctx.Value("location").(db.Location)
	uid := chi.URLParam(request, "evse_id")
	dto := EvseDto{}

	if err := json.NewDecoder(request.Body).Decode(&dto); err != nil {
		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	evse := r.ReplaceEvse(ctx, location.ID, uid, &dto)

	if evse != nil {
		if dto.Capabilities != nil || dto.Status != nil {
			r.updateLocationAvailability(ctx, evse.ID)
		}

		err := r.Repository.UpdateLocationLastUpdated(ctx, db.UpdateLocationLastUpdatedParams{
			ID:          location.ID,
			LastUpdated: evse.LastUpdated,
		})

		if err != nil {
			render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
			return
		}
	}

	render.Render(rw, request, transportation.OcpiSuccess(nil))
}
