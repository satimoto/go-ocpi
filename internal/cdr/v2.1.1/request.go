package cdr

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
)

func (r *CdrResolver) GetCdr(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cdr := ctx.Value("cdr").(db.Cdr)
	dto := r.CreateCdrDto(ctx, cdr)

	if err := render.Render(rw, request, ocpi.OCPISuccess(dto)); err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}
}

func (r *CdrResolver) PostCdr(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	dto, err := r.UnmarshalDto(request.Body)

	if err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}

	cdr := r.CreateCdr(ctx, dto)

	if cdr == nil {
		render.Render(rw, request, ocpi.OCPIErrorMissingParameters(nil))
	}

	location := fmt.Sprintf("%s/%s/%s/%s", os.Getenv("API_DOMAIN"), "2.1.1", "cdrs", cdr.Uid)
	rw.Header().Add("Location", location)
	render.Render(rw, request, ocpi.OCPISuccess(nil))
}
