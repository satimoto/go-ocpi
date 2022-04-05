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
	dto := r.CreateLocationDto(ctx, location)

	if err := render.Render(rw, request, ocpi.OCPISuccess(dto)); err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}
}

func (r *LocationResolver) UpdateLocation(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	countryCode := chi.URLParam(request, "country_code")
	partyID := chi.URLParam(request, "party_id")
	uid := chi.URLParam(request, "location_id")
	dto, err := r.UnmarshalDto(request.Body)

	if err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}

	location := r.ReplaceLocationWithIdentifiers(ctx, &countryCode, &partyID, uid, dto)

	if location == nil {
		render.Render(rw, request, ocpi.OCPIErrorMissingParameters(nil))
	}

	render.Render(rw, request, ocpi.OCPISuccess(nil))
}
