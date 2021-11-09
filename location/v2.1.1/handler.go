package location

import (
	"context"
	"encoding/json"
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

func (r *LocationResolver) routes() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/{country_code}/{party_id}", func(credentialRouter chi.Router) {
		credentialRouter.Use(rest.Credential(r.Repository))

		credentialRouter.Route("/{locationID}", func(locationRouter chi.Router) {
			locationRouter.Put("/", r.updateLocation)

			locationContextRouter := locationRouter.With(r.locationContext)
			locationContextRouter.Get("/", r.getLocation)
			locationContextRouter.Patch("/", r.updateLocation)

			locationContextRouter.Mount("/", r.EvseResolver.Routes())
		})
	})

	return router
}

func (r *LocationResolver) locationContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if locationID := chi.URLParam(request, "locationID"); locationID != "" {
			location, err := r.Repository.GetLocationByUid(ctx, locationID)

			if err != nil {
				render.Render(rw, request, rest.OCPIErrorUnknownResource(nil))
				return
			}

			ctx = context.WithValue(ctx, "location", location)
		}

		next.ServeHTTP(rw, request.WithContext(ctx))
	})
}

func (r *LocationResolver) getLocation(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	location := ctx.Value("location").(db.Location)
	if err := render.Render(rw, request, r.CreateLocationPayload(ctx, location)); err != nil {
		render.Render(rw, request, rest.OCPIServerError(nil, err.Error()))
	}
}

func (r *LocationResolver) updateLocation(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	uid := chi.URLParam(request, "locationID")
	payload := LocationPayload{}

	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		render.Render(rw, request, rest.OCPIServerError(nil, err.Error()))
	}

	location := r.ReplaceLocation(ctx, uid, &payload)

	if location == nil {
		render.Render(rw, request, rest.OCPIErrorMissingParameters(nil))
	}

	render.Render(rw, request, rest.OCPISuccess(nil))
}
