package session

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi-api/internal/middleware"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *SessionResolver) GetSession(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	session := ctx.Value("session").(db.Session)
	dto := r.CreateSessionDto(ctx, session)

	if err := render.Render(rw, request, transportation.OcpiSuccess(dto)); err != nil {
		util.LogOnError("OCPI175", "Error rendering response", err)
		util.LogHttpRequest("OCPI175", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}

func (r *SessionResolver) UpdateSession(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred := middleware.GetCredential(ctx)
	countryCode := chi.URLParam(request, "country_code")
	partyID := chi.URLParam(request, "party_id")
	uid := chi.URLParam(request, "session_id")
	dto, err := r.UnmarshalPushDto(request.Body)

	if err != nil {
		util.LogOnError("OCPI176", "Error unmarshalling request", err)
		util.LogHttpRequest("OCPI176", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	session := r.ReplaceSessionByIdentifier(ctx, *cred, &countryCode, &partyID, uid, dto)

	if session == nil {
		log.Print("OCPI177", "Error replacing cdr")
		util.LogHttpRequest("OCPI177", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiErrorMissingParameters(nil))
		return
	}

	render.Render(rw, request, transportation.OcpiSuccess(nil))
}
