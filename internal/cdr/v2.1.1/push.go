package cdr

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/middleware"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/version"
)

func (r *CdrResolver) GetCdr(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cdr := ctx.Value("cdr").(db.Cdr)
	dto := r.CreateCdrDto(ctx, cdr)

	if err := render.Render(rw, request, transportation.OcpiSuccess(dto)); err != nil {
		util.LogOnError("OCPI033", "Error rendering response", err)
		util.LogHttpRequest("OCPI033", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}

func (r *CdrResolver) PostCdr(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred := middleware.GetCredential(ctx)
	dto, err := r.UnmarshalPushDto(request.Body)

	if err != nil {
		util.LogOnError("OCPI034", "Error unmarshalling request", err)
		util.LogHttpRequest("OCPI034", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	cdr := r.ReplaceCdr(ctx, *cred, *dto.ID, dto)

	if cdr == nil {
		log.Print("OCPI035", "Error replacing cdr")
		util.LogHttpRequest("OCPI035", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiErrorMissingParameters(nil))
		return
	}

	location := fmt.Sprintf("%s/%s/%s/%s", os.Getenv("API_DOMAIN"), version.VERSION_2_1_1, "cdrs", cdr.Uid)
	rw.Header().Add("Location", location)
	render.Render(rw, request, transportation.OcpiSuccess(nil))
}
