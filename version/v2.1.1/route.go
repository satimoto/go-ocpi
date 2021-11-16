package version

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/credential"
	"github.com/satimoto/go-ocpi-api/rest"
)

func New(repositoryService *db.RepositoryService) *chi.Mux {
	r := NewResolver(repositoryService)

	return r.routes()
}

func (r *VersionDetailResolver) routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(credential.CredentialContextByToken(r.Repository))
	router.Get("/", r.getVersionDetail)

	return router
}

func (r *VersionDetailResolver) getVersionDetail(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	payload := r.CreateVersionDetailPayload(ctx)

	if err := render.Render(rw, request, rest.OCPISuccess(payload)); err != nil {
		render.Render(rw, request, rest.OCPIServerError(nil, err.Error()))
	}
}