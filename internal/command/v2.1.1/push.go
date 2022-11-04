package command

import (
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *CommandResolver) PostCommandReservationResponse(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	command := ctx.Value("command").(db.CommandReservation)
	dto, err := r.UnmarshalPushDto(request.Body)

	if err != nil {
		metrics.RecordError("OCPI071", "Error unmarshalling request", err)
		util.LogHttpRequest("OCPI071", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	r.UpdateCommandReservation(ctx, command, dto)

	if err := render.Render(rw, request, transportation.OcpiSuccess(nil)); err != nil {
		log.Print("OCPI072", "Error updating reservation command")
		util.LogHttpRequest("OCPI072", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}

func (r *CommandResolver) PostCommandStartResponse(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	command := ctx.Value("command").(db.CommandStart)
	dto, err := r.UnmarshalPushDto(request.Body)

	if err != nil {
		metrics.RecordError("OCPI073", "Error unmarshalling request", err)
		util.LogHttpRequest("OCPI073", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	r.UpdateCommandStart(ctx, command, dto)

	if err := render.Render(rw, request, transportation.OcpiSuccess(nil)); err != nil {
		log.Print("OCPI074", "Error updating start command")
		util.LogHttpRequest("OCPI074", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}

func (r *CommandResolver) PostCommandStopResponse(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	command := ctx.Value("command").(db.CommandStop)
	dto, err := r.UnmarshalPushDto(request.Body)

	if err != nil {
		metrics.RecordError("OCPI075", "Error unmarshalling request", err)
		util.LogHttpRequest("OCPI075", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	r.UpdateCommandStop(ctx, command, dto)

	if err := render.Render(rw, request, transportation.OcpiSuccess(nil)); err != nil {
		log.Print("OCPI076", "Error updating stop command")
		util.LogHttpRequest("OCPI076", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}

func (r *CommandResolver) PostCommandUnlockResponse(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	command := ctx.Value("command").(db.CommandUnlock)
	dto, err := r.UnmarshalPushDto(request.Body)

	if err != nil {
		metrics.RecordError("OCPI077", "Error unmarshalling request", err)
		util.LogHttpRequest("OCPI077", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	r.UpdateCommandUnlock(ctx, command, dto)

	if err := render.Render(rw, request, transportation.OcpiSuccess(nil)); err != nil {
		log.Print("OCPI078", "Error updating unlock command")
		util.LogHttpRequest("OCPI078", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}
