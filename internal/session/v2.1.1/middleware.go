package session

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *SessionResolver) SessionContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if sessionID := chi.URLParam(request, "session_id"); sessionID != "" {
			session, err := r.Repository.GetSessionByUid(ctx, sessionID)

			if err == nil {
				ctx = context.WithValue(ctx, "session", session)
				next.ServeHTTP(rw, request.WithContext(ctx))
				return
			}
		}

		render.Render(rw, request, transportation.OcpiErrorUnknownResource(nil))
	})
}
