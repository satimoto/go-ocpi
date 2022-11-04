package version

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/satimoto/go-datastore/pkg/util"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *VersionResolver) GetVersions(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	dto := r.CreateVersionListDto(ctx)

	if err := render.Render(rw, request, transportation.OcpiSuccess(dto)); err != nil {
		metrics.RecordError("OCPI215", "Error rendering response", err)
		util.LogHttpRequest("OCPI215", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}
