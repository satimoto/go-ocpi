package command

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
)

func (r *CommandResolver) PostCommandReservationResponse(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	command := ctx.Value("command").(db.CommandReservation)
	commandResponseDto, err := r.UnmarshalCommandResponseDto(request.Body)

	if err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}

	r.UpdateCommandReservation(ctx, command, commandResponseDto)

	if err := render.Render(rw, request, ocpi.OCPISuccess(nil)); err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}
}

func (r *CommandResolver) PostCommandStartResponse(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	command := ctx.Value("command").(db.CommandStart)
	commandResponseDto, err := r.UnmarshalCommandResponseDto(request.Body)

	if err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}

	r.UpdateCommandStart(ctx, command, commandResponseDto)

	if err := render.Render(rw, request, ocpi.OCPISuccess(nil)); err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}
}

func (r *CommandResolver) PostCommandStopResponse(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	command := ctx.Value("command").(db.CommandStop)
	commandResponseDto, err := r.UnmarshalCommandResponseDto(request.Body)

	if err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}

	r.UpdateCommandStop(ctx, command, commandResponseDto)

	if err := render.Render(rw, request, ocpi.OCPISuccess(nil)); err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}
}

func (r *CommandResolver) PostCommandUnlockResponse(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	command := ctx.Value("command").(db.CommandUnlock)
	commandResponseDto, err := r.UnmarshalCommandResponseDto(request.Body)

	if err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}

	r.UpdateCommandUnlock(ctx, command, commandResponseDto)

	if err := render.Render(rw, request, ocpi.OCPISuccess(nil)); err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}
}
