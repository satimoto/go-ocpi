package location

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/credential"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *LocationResolver) GetLocation(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	location := ctx.Value("location").(db.Location)
	dto := r.CreateLocationDto(ctx, location)

	if err := render.Render(rw, request, transportation.OCPISuccess(dto)); err != nil {
		render.Render(rw, request, transportation.OCPIServerError(nil, err.Error()))
	}
}

func (r *LocationResolver) UpdateLocation(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred := credential.GetCredential(ctx)
	countryCode := chi.URLParam(request, "country_code")
	partyID := chi.URLParam(request, "party_id")
	uid := chi.URLParam(request, "location_id")
	dto, err := r.UnmarshalPushDto(request.Body)

	if err != nil {
		render.Render(rw, request, transportation.OCPIServerError(nil, err.Error()))
		return
	}

	location := r.ReplaceLocationByIdentifier(ctx, *cred, &countryCode, &partyID, uid, dto)

	if location == nil {
		render.Render(rw, request, transportation.OCPIErrorMissingParameters(nil))
		return
	}

	render.Render(rw, request, transportation.OCPISuccess(nil))
}
