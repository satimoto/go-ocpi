package tokenauthorization

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *TokenAuthorizationResolver) AuthorizeToken(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	token := ctx.Value("token").(db.Token)
	locationReferencesDto, err := r.UnmarshalLocationReferencesDto(request.Body)

	if err != nil {
		render.Render(rw, request, transportation.OCPIServerError(nil, err.Error()))
		return
	}

	tokenAuthorization := r.CreateTokenAuthorization(ctx, token, locationReferencesDto)
	dto := r.CreateAuthorizationInfoDto(ctx, token, tokenAuthorization, locationReferencesDto, nil)

	if err := render.Render(rw, request, transportation.OCPISuccess(dto)); err != nil {
		render.Render(rw, request, transportation.OCPIServerError(nil, err.Error()))
	}
}
