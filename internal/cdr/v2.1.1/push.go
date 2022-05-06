package cdr

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/middleware"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *CdrResolver) GetCdr(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cdr := ctx.Value("cdr").(db.Cdr)
	dto := r.CreateCdrDto(ctx, cdr)

	if err := render.Render(rw, request, transportation.OcpiSuccess(dto)); err != nil {
		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}

func (r *CdrResolver) PostCdr(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred := middleware.GetCredential(ctx)
	dto, err := r.UnmarshalPushDto(request.Body)

	if err != nil {
		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	cdr := r.ReplaceCdr(ctx, *cred, dto)

	if cdr == nil {
		render.Render(rw, request, transportation.OcpiErrorMissingParameters(nil))
		return
	}

	location := fmt.Sprintf("%s/%s/%s/%s", os.Getenv("API_DOMAIN"), "2.1.1", "cdrs", cdr.Uid)
	rw.Header().Add("Location", location)
	render.Render(rw, request, transportation.OcpiSuccess(nil))
}
