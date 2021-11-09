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

func (r *VersionResolver) routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", r.getVersions)

	return router
}

func (r *VersionResolver) getVersions(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	if err := render.RenderList(rw, request, r.CreateVersionListPayload(ctx)); err != nil {
		render.Render(rw, request, rest.OCPIServerError(nil, err.Error()))
	}
}