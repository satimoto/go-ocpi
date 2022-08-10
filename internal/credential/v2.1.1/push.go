package credential

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/middleware"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *CredentialResolver) DeleteCredential(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred, err := middleware.GetCredentialByToken(r.Repository, ctx, request)

	if err != nil || !cred.ClientToken.Valid {
		util.LogOnError("OCPI085", "Error retrieving credential", err)
		util.LogHttpRequest("OCPI085", request.URL.String(), request, true)

		http.Error(rw, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	updateCredentialParams := param.NewUpdateCredentialParams(cred)
	updateCredentialParams.ClientToken = util.SqlNullString(nil)
	updateCredentialParams.LastUpdated = time.Now()

	cred, err = r.Repository.UpdateCredential(ctx, updateCredentialParams)

	if err != nil {
		log.Print("OCPI086", "Error updating credential")
		util.LogHttpRequest("OCPI086", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	render.Render(rw, request, transportation.OcpiSuccess(nil))
}

func (r *CredentialResolver) GetCredential(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred := ctx.Value("credential").(db.Credential)
	dto := r.CreateCredentialDto(ctx, cred)

	if err := render.Render(rw, request, transportation.OcpiSuccess(dto)); err != nil {
		util.LogOnError("OCPI087", "Error rendering response", err)
		util.LogHttpRequest("OCPI087", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}

func (r *CredentialResolver) UpdateCredential(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred, err := middleware.GetCredentialByToken(r.Repository, ctx, request)
	dto := CredentialDto{}

	if err != nil || !cred.ClientToken.Valid {
		util.LogOnError("OCPI088", "Error retrieving credential", err)
		util.LogHttpRequest("OCPI088", request.URL.String(), request, true)

		http.Error(rw, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(request.Body).Decode(&dto); err != nil {
		util.LogOnError("OCPI089", "Error unmarshalling request", err)
		util.LogHttpRequest("OCPI089", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	c, err := r.ReplaceCredential(ctx, cred, &dto)

	if err != nil {
		util.LogOnError("OCPI090", "Error replacing credential", err)
		util.LogHttpRequest("OCPI090", request.URL.String(), request, true)

		errResponse := err.(*transportation.OcpiResponse)
		render.Render(rw, request, errResponse)
		return
	}

	credentialDto := r.CreateCredentialDto(ctx, *c)
	render.Render(rw, request, transportation.OcpiSuccess(credentialDto))
}
