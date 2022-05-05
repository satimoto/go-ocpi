package versiondetail

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *VersionDetailResolver) GetVersionDetail(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	dto := r.CreateVersionDetailDto(ctx)

	if err := render.Render(rw, request, transportation.OCPISuccess(dto)); err != nil {
		render.Render(rw, request, transportation.OCPIServerError(nil, err.Error()))
	}
}
