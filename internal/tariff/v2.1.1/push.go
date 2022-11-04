package tariff

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/middleware"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *TariffResolver) DeleteTariff(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	tariff := ctx.Value("tariff").(db.Tariff)
	err := r.Repository.DeleteTariffByUid(ctx, tariff.Uid)

	if err != nil {
		metrics.RecordError("OCPI187", "Error deleting tariff", err)
		util.LogHttpRequest("OCPI187", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}

	render.Render(rw, request, transportation.OcpiSuccess(nil))
}

func (r *TariffResolver) GetTariff(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	tariff := ctx.Value("tariff").(db.Tariff)
	dto := r.CreateTariffDto(ctx, tariff)

	if err := render.Render(rw, request, transportation.OcpiSuccess(dto)); err != nil {
		metrics.RecordError("OCPI188", "Error rendering response", err)
		util.LogHttpRequest("OCPI188", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}

func (r *TariffResolver) UpdateTariff(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred := middleware.GetCredential(ctx)
	countryCode := chi.URLParam(request, "country_code")
	partyID := chi.URLParam(request, "party_id")
	uid := chi.URLParam(request, "tariff_id")
	dto, err := r.UnmarshalPushDto(request.Body)

	if err != nil {
		metrics.RecordError("OCPI189", "Error unmarshalling request", err)
		util.LogHttpRequest("OCPI189", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	tariff := r.ReplaceTariffByIdentifier(ctx, *cred, &countryCode, &partyID, uid, nil, dto)

	if tariff == nil {
		log.Print("OCPI190", "Error replacing cdr")
		util.LogHttpRequest("OCPI190", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiErrorMissingParameters(nil))
		return
	}

	render.Render(rw, request, transportation.OcpiSuccess(nil))
}
