package versiondetail

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/satimoto/go-ocpi-api/internal/ocpi"
)

func (r *VersionDetailResolver) GetVersionDetail(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	dto := r.CreateVersionDetailDto(ctx)

	if err := render.Render(rw, request, ocpi.OCPISuccess(dto)); err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}
}
