package tokenauthorization

import (
	"io"
	"net/http"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *TokenAuthorizationResolver) AuthorizeToken(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	token := ctx.Value("token").(db.Token)
	locationReferencesDto, err := r.UnmarshalLocationReferencesDto(request.Body)

	if err != nil && err != io.EOF {
		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	// TODO: we should reject an authorization to an unknown location/evse/connector
	tokenAuthorization, err := r.CreateTokenAuthorization(ctx, token, locationReferencesDto)

	if err != nil {
		render.Render(rw, request, transportation.OcpiErrorNotEnoughInformation(nil))
		return
	}

	dto := r.CreateAuthorizationInfoDto(ctx, token, tokenAuthorization, locationReferencesDto, nil)

	if err := render.Render(rw, request, transportation.OcpiSuccess(dto)); err != nil {
		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}
