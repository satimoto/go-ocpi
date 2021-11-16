package evse

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/rest"
)

func (r *EvseResolver) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/{evseId}", func(evseRouter chi.Router) {
		evseRouter.Put("/", r.updateEvse)

		evseContextRouter := evseRouter.With(r.evseContext)
		evseContextRouter.Get("/", r.getEvse)
		evseContextRouter.Patch("/", r.updateEvse)

		evseContextRouter.Mount("/", r.ConnectorResolver.Routes())
	})

	return router
}

func (r *EvseResolver) evseContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if evseID := chi.URLParam(request, "evseID"); evseID != "" {
			evse, err := r.Repository.GetEvseByUid(ctx, evseID)

			if err == nil {
				ctx = context.WithValue(ctx, "evse", evse)
				next.ServeHTTP(rw, request.WithContext(ctx))
				return
			}
		}

		render.Render(rw, request, rest.OCPIErrorUnknownResource(nil))
	})
}

func (r *EvseResolver) getEvse(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	evse := ctx.Value("evse").(db.Evse)
	payload := r.CreateEvsePayload(ctx, evse)

	if err := render.Render(rw, request, rest.OCPISuccess(payload)); err != nil {
		render.Render(rw, request, rest.OCPIServerError(nil, err.Error()))
	}
}

func (r *EvseResolver) updateEvse(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	location := ctx.Value("location").(db.Location)
	uid := chi.URLParam(request, "evseID")
	payload := EvsePayload{}

	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		render.Render(rw, request, rest.OCPIServerError(nil, err.Error()))
	}

	evse := r.ReplaceEvse(ctx, location.ID, uid, &payload)

	err := r.Repository.UpdateLocationLastUpdated(ctx, db.UpdateLocationLastUpdatedParams{
		ID:          location.ID,
		LastUpdated: evse.LastUpdated,
	})

	if err != nil {
		render.Render(rw, request, rest.OCPIServerError(nil, err.Error()))
	}

	render.Render(rw, request, rest.OCPISuccess(nil))
}
