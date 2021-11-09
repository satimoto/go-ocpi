package version

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/rest"
)

func New(repositoryService *db.RepositoryService) *chi.Mux {
	r := NewResolver(repositoryService)

	return r.routes()
}

func (r *VersionDetailResolver) routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", r.getVersionDetail)

	return router
}

func (r *VersionDetailResolver) getVersionDetail(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	if err := render.Render(rw, request, r.CreateVersionDetailPayload(ctx)); err != nil {
		render.Render(rw, request, rest.OCPIServerError(nil, err.Error()))
	}
}