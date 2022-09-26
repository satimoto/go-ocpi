package location

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/middleware"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *LocationResolver) GetLocation(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	location := ctx.Value("location").(db.Location)
	dto := r.CreateLocationDto(ctx, location)

	if err := render.Render(rw, request, transportation.OcpiSuccess(dto)); err != nil {
		util.LogOnError("OCPI130", "Error rendering response", err)
		util.LogHttpRequest("OCPI130", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}

func (r *LocationResolver) UpdateLocation(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred := middleware.GetCredential(ctx)
	countryCode := chi.URLParam(request, "country_code")
	partyID := chi.URLParam(request, "party_id")
	uid := chi.URLParam(request, "location_id")
	dto, err := r.UnmarshalPushDto(request.Body)

	if err != nil {
		util.LogOnError("OCPI131", "Error unmarshalling request", err)
		util.LogHttpRequest("OCPI131", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	location := r.ReplaceLocationByIdentifier(ctx, *cred, &countryCode, &partyID, uid, dto)

	if location == nil {
		util.LogOnError("OCPI132", "Error replacing location", err)
		util.LogHttpRequest("OCPI132", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiErrorMissingParameters(nil))
		return
	}

	render.Render(rw, request, transportation.OcpiSuccess(nil))
}
