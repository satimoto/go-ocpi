package evse

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
)

func (r *EvseResolver) GetEvse(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	evse := ctx.Value("evse").(db.Evse)
	payload := r.CreateEvsePayload(ctx, evse)

	if err := render.Render(rw, request, ocpi.OCPISuccess(payload)); err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}
}

func (r *EvseResolver) UpdateEvse(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	location := ctx.Value("location").(db.Location)
	uid := chi.URLParam(request, "evse_id")
	payload := EvsePayload{}

	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}

	evse := r.ReplaceEvse(ctx, location.ID, uid, &payload)

	err := r.Repository.UpdateLocationLastUpdated(ctx, db.UpdateLocationLastUpdatedParams{
		ID:          location.ID,
		LastUpdated: evse.LastUpdated,
	})

	if err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}

	render.Render(rw, request, ocpi.OCPISuccess(nil))
}
