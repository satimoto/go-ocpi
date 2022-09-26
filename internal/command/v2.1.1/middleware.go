package command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *CommandResolver) CommandReservationContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if commandID := chi.URLParam(request, "command_id"); commandID != "" {
			if id, err := strconv.ParseInt(commandID, 10, 64); err == nil {
				command, err := r.Repository.GetCommandReservation(ctx, id)

				if err == nil {
					ctx = context.WithValue(ctx, "command", command)
					next.ServeHTTP(rw, request.WithContext(ctx))
					return
				}
			}
		}

		render.Render(rw, request, transportation.OcpiErrorUnknownResource(nil))
	})
}

func (r *CommandResolver) CommandStartContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if commandID := chi.URLParam(request, "command_id"); commandID != "" {
			if id, err := strconv.ParseInt(commandID, 10, 64); err == nil {
				command, err := r.Repository.GetCommandStart(ctx, id)

				if err == nil {
					ctx = context.WithValue(ctx, "command", command)
					next.ServeHTTP(rw, request.WithContext(ctx))
					return
				}
			}
		}

		render.Render(rw, request, transportation.OcpiErrorUnknownResource(nil))
	})
}

func (r *CommandResolver) CommandStopContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if commandID := chi.URLParam(request, "command_id"); commandID != "" {
			if id, err := strconv.ParseInt(commandID, 10, 64); err == nil {
				command, err := r.Repository.GetCommandStop(ctx, id)

				if err == nil {
					ctx = context.WithValue(ctx, "command", command)
					next.ServeHTTP(rw, request.WithContext(ctx))
					return
				}
			}
		}

		render.Render(rw, request, transportation.OcpiErrorUnknownResource(nil))
	})
}

func (r *CommandResolver) CommandUnlockContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if commandID := chi.URLParam(request, "command_id"); commandID != "" {
			if id, err := strconv.ParseInt(commandID, 10, 64); err == nil {
				command, err := r.Repository.GetCommandUnlock(ctx, id)

				if err == nil {
					ctx = context.WithValue(ctx, "command", command)
					next.ServeHTTP(rw, request.WithContext(ctx))
					return
				}
			}
		}

		render.Render(rw, request, transportation.OcpiErrorUnknownResource(nil))
	})
}
