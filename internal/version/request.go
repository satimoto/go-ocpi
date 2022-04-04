package version

import (
	"net/http"

	"github.com/go-chi/render"

	//credential "github.com/satimoto/go-ocpi-api/internal/credential"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
)

func (r *VersionResolver) GetVersions(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	dto := r.CreateVersionListDto(ctx)

	if err := render.Render(rw, request, ocpi.OCPISuccess(dto)); err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}
}
