package session

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/middleware"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *SessionResolver) GetSession(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	session := ctx.Value("session").(db.Session)
	dto := r.CreateSessionDto(ctx, session)

	if err := render.Render(rw, request, transportation.OcpiSuccess(dto)); err != nil {
		metrics.RecordError("OCPI175", "Error rendering response", err)
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
		metrics.RecordError("OCPI176", "Error unmarshaling request", err)
		util.LogHttpRequest("OCPI176", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	session := r.ReplaceSessionByIdentifier(ctx, *cred, &countryCode, &partyID, uid, dto)

	if session == nil {
		log.Print("OCPI177", "Error replacing session")
		util.LogHttpRequest("OCPI177", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiErrorMissingParameters(nil))
		return
	}

	render.Render(rw, request, transportation.OcpiSuccess(nil))
}
