package credential

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/credential"
	"github.com/satimoto/go-ocpi-api/rest"
	"github.com/satimoto/go-ocpi-api/util"
)

func New(repositoryService *db.RepositoryService) *chi.Mux {
	r := NewResolver(repositoryService)

	return r.routes()
}

func (r *CredentialResolver) routes() *chi.Mux {
	router := chi.NewRouter()

	router.Delete("/", r.deleteCredential)
	router.Post("/", r.postCredential)
	router.Put("/", r.putCredential)

	credentialContextRouter := router.With(credential.CredentialContextByToken(r.Repository))
	credentialContextRouter.Get("/", r.getCredential)

	return router
}

func (r *CredentialResolver) deleteCredential(rw http.ResponseWriter, request *http.Request) {
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
		render.Render(rw, request, rest.OCPIServerError(nil, err.Error()))
	}

	render.Render(rw, request, rest.OCPISuccess(nil))
}

func (r *CredentialResolver) getCredential(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred := ctx.Value("credential").(db.Credential)
	payload := r.CreateCredentialPayload(ctx, cred)

	if err := render.Render(rw, request, rest.OCPISuccess(payload)); err != nil {
		render.Render(rw, request, rest.OCPIServerError(nil, err.Error()))
	}
}

func (r *CredentialResolver) postCredential(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred, err := credential.GetCredentialByToken(r.Repository, ctx, request)

	if err == nil && len(cred.ClientToken.String) == 0 {
		r.updateCredential(ctx, cred, rw, request)
		return
	}

	http.Error(rw, http.StatusText(405), 405)
}

func (r *CredentialResolver) putCredential(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred, err := credential.GetCredentialByToken(r.Repository, ctx, request)

	if err == nil && len(cred.ClientToken.String) > 0 {
		r.updateCredential(ctx, cred, rw, request)
		return
	}

	http.Error(rw, http.StatusText(405), 405)
}

func (r *CredentialResolver) updateCredential(ctx context.Context, cred db.Credential, rw http.ResponseWriter, request *http.Request) {
	payload := CredentialPayload{}

	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		render.Render(rw, request, rest.OCPIServerError(nil, err.Error()))
	}

	c, err := r.ReplaceCredential(ctx, cred, &payload)

	if err != nil {
		errResponse := err.(*rest.OCPIResponse)
		render.Render(rw, request, errResponse)
	}

	credentialPayload := r.CreateCredentialPayload(ctx, *c)
	render.Render(rw, request, rest.OCPISuccess(credentialPayload))
}
