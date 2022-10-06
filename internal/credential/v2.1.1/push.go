package credential

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/param"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/middleware"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/util"
)

func (r *CredentialResolver) DeleteCredential(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred, err := middleware.GetCredentialByToken(r.Repository, ctx, request)

	if err != nil || !cred.ClientToken.Valid {
		dbUtil.LogOnError("OCPI085", "Error retrieving credential", err)
		dbUtil.LogHttpRequest("OCPI085", request.URL.String(), request, true)

		http.Error(rw, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	updateCredentialParams := param.NewUpdateCredentialParams(cred)
	updateCredentialParams.ClientToken = dbUtil.SqlNullString(nil)
	updateCredentialParams.LastUpdated = util.NewTimeUTC()

	cred, err = r.Repository.UpdateCredential(ctx, updateCredentialParams)

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
		dbUtil.LogOnError("OCPI087", "Error rendering response", err)
		dbUtil.LogHttpRequest("OCPI087", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}

func (r *CredentialResolver) UpdateCredential(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred, err := middleware.GetCredentialByToken(r.Repository, ctx, request)
	credentialDto := dto.CredentialDto{}

	if err != nil || !cred.ClientToken.Valid {
		dbUtil.LogOnError("OCPI088", "Error retrieving credential", err)
		dbUtil.LogHttpRequest("OCPI088", request.URL.String(), request, true)

		http.Error(rw, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(request.Body).Decode(&credentialDto); err != nil {
		dbUtil.LogOnError("OCPI089", "Error unmarshalling request", err)
		dbUtil.LogHttpRequest("OCPI089", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	c, err := r.ReplaceCredential(ctx, cred, &credentialDto)

	if err != nil {
		dbUtil.LogOnError("OCPI090", "Error replacing credential", err)
		dbUtil.LogHttpRequest("OCPI090", request.URL.String(), request, true)

		errResponse := err.(*transportation.OcpiResponse)
		render.Render(rw, request, errResponse)
		return
	}

	createCredentialDto := r.CreateCredentialDto(ctx, *c)
	render.Render(rw, request, transportation.OcpiSuccess(createCredentialDto))
}
