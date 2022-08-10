package versiondetail

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *VersionDetailResolver) GetVersionDetail(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	dto := r.CreateVersionDetailDto(ctx)

	if err := render.Render(rw, request, transportation.OcpiSuccess(dto)); err != nil {
		util.LogOnError("OCPI220", "Error rendering response", err)
		util.LogHttpRequest("OCPI220", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}
