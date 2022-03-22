package location

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
)


func (r *LocationResolver) GetLocation(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	location := ctx.Value("location").(db.Location)
	payload := r.CreateLocationPayload(ctx, location)

	if err := render.Render(rw, request, ocpi.OCPISuccess(payload)); err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}
}

func (r *LocationResolver) UpdateLocation(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	uid := chi.URLParam(request, "location_id")
	payload, err := r.UnmarshalPayload(request.Body)

	if err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}

	location := r.ReplaceLocation(ctx, uid, payload)

	if location == nil {
		render.Render(rw, request, ocpi.OCPIErrorMissingParameters(nil))
	}

	render.Render(rw, request, ocpi.OCPISuccess(nil))
}
