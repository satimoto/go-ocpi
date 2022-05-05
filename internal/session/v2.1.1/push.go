package session

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/credential"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *SessionResolver) GetSession(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	session := ctx.Value("session").(db.Session)
	dto := r.CreateSessionDto(ctx, session)

	if err := render.Render(rw, request, transportation.OCPISuccess(dto)); err != nil {
		render.Render(rw, request, transportation.OCPIServerError(nil, err.Error()))
	}
}

func (r *SessionResolver) UpdateSession(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred := credential.GetCredential(ctx)
	countryCode := chi.URLParam(request, "country_code")
	partyID := chi.URLParam(request, "party_id")
	uid := chi.URLParam(request, "session_id")
	dto, err := r.UnmarshalPushDto(request.Body)

	if err != nil {
		render.Render(rw, request, transportation.OCPIServerError(nil, err.Error()))
		return
	}

	session := r.ReplaceSessionByIdentifier(ctx, *cred, &countryCode, &partyID, uid, dto)

	if session == nil {
		render.Render(rw, request, transportation.OCPIErrorMissingParameters(nil))
		return
	}

	render.Render(rw, request, transportation.OCPISuccess(nil))
}
