package credential

import (
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/param"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/middleware"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/util"
)

func (r *CredentialResolver) DeleteCredential(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred := middleware.GetCredential(ctx)

	if !cred.ClientToken.Valid {
		log.Print("OCPI085", "Error credentials are not registered")
		dbUtil.LogHttpRequest("OCPI085", request.URL.String(), request, true)

		http.Error(rw, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	updateCredentialParams := param.NewUpdateCredentialParams(*cred)
	updateCredentialParams.ClientToken = dbUtil.SqlNullString(nil)
	updateCredentialParams.LastUpdated = util.NewTimeUTC()

	_, err := r.Repository.UpdateCredential(ctx, updateCredentialParams)

	if err != nil {
		log.Print("OCPI086", "Error updating credential")
		dbUtil.LogHttpRequest("OCPI086", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	render.Render(rw, request, transportation.OcpiSuccess(nil))
}

func (r *CredentialResolver) GetCredential(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred := middleware.GetCredential(ctx)
	dto := r.CreateCredentialDto(ctx, *cred)

	if err := render.Render(rw, request, transportation.OcpiSuccess(dto)); err != nil {
		metrics.RecordError("OCPI087", "Error rendering response", err)
		dbUtil.LogHttpRequest("OCPI087", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}

func (r *CredentialResolver) UpdateCredential(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred := middleware.GetCredential(ctx)

	if request.Method == http.MethodPost && cred.ClientToken.Valid {
		log.Printf("OCPI088: Error credentials are registered")
		dbUtil.LogHttpRequest("OCPI088", request.URL.String(), request, true)

		http.Error(rw, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if request.Method == http.MethodPut && !cred.ClientToken.Valid {
		log.Printf("OCPI019: Error credentials are not registered")
		dbUtil.LogHttpRequest("OCPI019", request.URL.String(), request, true)

		http.Error(rw, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	dto, err := r.UnmarshalPushDto(request.Body)

	if err != nil {
		metrics.RecordError("OCPI089", "Error unmarshalling request", err)
		dbUtil.LogHttpRequest("OCPI089", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	c, err := r.ReplaceCredential(ctx, *cred, dto)

	if err != nil {
		metrics.RecordError("OCPI090", "Error replacing credential", err)
		dbUtil.LogHttpRequest("OCPI090", request.URL.String(), request, true)

		errResponse := err.(*transportation.OcpiResponse)
		render.Render(rw, request, errResponse)
		return
	}

	createCredentialDto := r.CreateCredentialDto(ctx, *c)
	render.Render(rw, request, transportation.OcpiSuccess(createCredentialDto))
}
