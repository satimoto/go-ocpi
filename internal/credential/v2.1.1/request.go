package credential

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/credential"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
	"github.com/satimoto/go-ocpi-api/internal/util"
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
	payload := r.CreateCredentialPayload(ctx, cred)

	if err := render.Render(rw, request, ocpi.OCPISuccess(payload)); err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}
}

func (r *CredentialResolver) UpdateCredential(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred, err := credential.GetCredentialByToken(r.Repository, ctx, request)
	payload := CredentialPayload{}

	if err == nil && len(cred.ClientToken.String) == 0 {
		if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
			render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
		}

		c, err := r.ReplaceCredential(ctx, cred, &payload)

		if err != nil {
			errResponse := err.(*ocpi.OCPIResponse)
			render.Render(rw, request, errResponse)
		}

		credentialPayload := r.CreateCredentialPayload(ctx, *c)
		render.Render(rw, request, ocpi.OCPISuccess(credentialPayload))
		return
	}

	http.Error(rw, http.StatusText(405), 405)
}
