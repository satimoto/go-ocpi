package tariff

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi-api/internal/middleware"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *TariffResolver) DeleteTariff(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	tariff := ctx.Value("tariff").(db.Tariff)
	err := r.Repository.DeleteTariffByUid(ctx, tariff.Uid)

	if err != nil {
		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}

	render.Render(rw, request, transportation.OcpiSuccess(nil))
}

func (r *TariffResolver) GetTariff(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	tariff := ctx.Value("tariff").(db.Tariff)
	dto := r.CreateTariffDto(ctx, tariff)

	if err := render.Render(rw, request, transportation.OcpiSuccess(dto)); err != nil {
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
		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	tariff := r.ReplaceTariffByIdentifier(ctx, *cred, &countryCode, &partyID, uid, nil, dto)

	if tariff == nil {
		render.Render(rw, request, transportation.OcpiErrorMissingParameters(nil))
		return
	}

	render.Render(rw, request, transportation.OcpiSuccess(nil))
}
