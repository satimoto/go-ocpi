package credential

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
	"github.com/satimoto/go-ocpi-api/internal/credential"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
)

func (r *CredentialResolver) DeleteCredential(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred, err := credential.GetCredentialByToken(r.Repository, ctx, request)

	if err != nil || !cred.ClientToken.Valid {
		http.Error(rw, http.StatusText(405), 405)
	}

	params := NewUpdateCredentialParams(cred)
	params.ClientToken = util.SqlNullString(nil)
	params.LastUpdated = time.Now()

	cred, err = r.Repository.UpdateCredential(ctx, db.UpdateCredentialParams{})

	if err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}

	render.Render(rw, request, ocpi.OCPISuccess(nil))
}

func (r *CredentialResolver) GetCredential(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred := ctx.Value("credential").(db.Credential)
	dto := r.CreateCredentialDto(ctx, cred)

	if err := render.Render(rw, request, ocpi.OCPISuccess(dto)); err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}
}

func (r *CredentialResolver) UpdateCredential(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred, err := credential.GetCredentialByToken(r.Repository, ctx, request)
	dto := CredentialDto{}

	if err == nil && len(cred.ClientToken.String) == 0 {
		if err := json.NewDecoder(request.Body).Decode(&dto); err != nil {
			render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
			return
		}

		c, err := r.ReplaceCredential(ctx, cred, &dto)

		if err != nil {
			errResponse := err.(*ocpi.OCPIResponse)
			render.Render(rw, request, errResponse)
			return
		}

		credentialDto := r.CreateCredentialDto(ctx, *c)
		render.Render(rw, request, ocpi.OCPISuccess(credentialDto))
		return
	}

	http.Error(rw, http.StatusText(405), 405)
}
