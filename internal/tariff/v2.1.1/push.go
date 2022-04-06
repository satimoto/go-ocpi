package tariff

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
)

func (r *TariffResolver) DeleteTariff(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	tariff := ctx.Value("tariff").(db.Tariff)
	err := r.Repository.DeleteTariffByUid(ctx, tariff.Uid)

	if err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}

	render.Render(rw, request, ocpi.OCPISuccess(nil))
}

func (r *TariffResolver) GetTariff(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	tariff := ctx.Value("tariff").(db.Tariff)
	dto := r.CreateTariffDto(ctx, tariff)

	if err := render.Render(rw, request, ocpi.OCPISuccess(dto)); err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}
}

func (r *TariffResolver) UpdateTariff(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	countryCode := chi.URLParam(request, "country_code")
	partyID := chi.URLParam(request, "party_id")
	uid := chi.URLParam(request, "tariff_id")
	dto, err := r.UnmarshalPushDto(request.Body)

	if err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}

	tariff := r.ReplaceTariffByIdentifier(ctx, &countryCode, &partyID, uid, nil, dto)

	if tariff == nil {
		render.Render(rw, request, ocpi.OCPIErrorMissingParameters(nil))
	}

	render.Render(rw, request, ocpi.OCPISuccess(nil))
}
