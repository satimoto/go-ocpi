package session

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
)

func (r *SessionResolver) GetSession(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	session := ctx.Value("session").(db.Session)
	dto := r.CreateSessionDto(ctx, session)

	if err := render.Render(rw, request, ocpi.OCPISuccess(dto)); err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}
}

func (r *SessionResolver) UpdateSession(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	uid := chi.URLParam(request, "session_id")
	dto, err := r.UnmarshalDto(request.Body)

	if err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}

	session := r.ReplaceSession(ctx, uid, dto)

	if session == nil {
		render.Render(rw, request, ocpi.OCPIErrorMissingParameters(nil))
	}

	render.Render(rw, request, ocpi.OCPISuccess(nil))
}
