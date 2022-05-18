package credential

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi-api/internal/credential"
	"github.com/satimoto/go-ocpi-api/internal/middleware"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *CredentialResolver) DeleteCredential(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred, err := middleware.GetCredentialByToken(r.Repository, ctx, request)

	if err != nil || !cred.ClientToken.Valid {
		http.Error(rw, http.StatusText(405), 405)
	}

	updateCredentialParams := credential.NewUpdateCredentialParams(cred)
	updateCredentialParams.ClientToken = util.SqlNullString(nil)
	updateCredentialParams.LastUpdated = time.Now()

	cred, err = r.Repository.UpdateCredential(ctx, updateCredentialParams)

	if err != nil {
		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}

	render.Render(rw, request, transportation.OcpiSuccess(nil))
}

func (r *CredentialResolver) GetCredential(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred := ctx.Value("credential").(db.Credential)
	dto := r.CreateCredentialDto(ctx, cred)

	if err := render.Render(rw, request, transportation.OcpiSuccess(dto)); err != nil {
		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}

func (r *CredentialResolver) UpdateCredential(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred, err := middleware.GetCredentialByToken(r.Repository, ctx, request)
	dto := CredentialDto{}

	if err == nil && len(cred.ClientToken.String) == 0 {
		if err := json.NewDecoder(request.Body).Decode(&dto); err != nil {
			render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
			return
		}

		c, err := r.ReplaceCredential(ctx, cred, &dto)

		if err != nil {
			errResponse := err.(*transportation.OcpiResponse)
			render.Render(rw, request, errResponse)
			return
		}

		credentialDto := r.CreateCredentialDto(ctx, *c)
		render.Render(rw, request, transportation.OcpiSuccess(credentialDto))
		return
	}

	http.Error(rw, http.StatusText(405), 405)
}
