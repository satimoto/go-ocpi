package tokenauthorization

import (
	"io"
	"net/http"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/middleware"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *TokenAuthorizationResolver) AuthorizeToken(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	cred := middleware.GetCredential(ctx)
	token := ctx.Value("token").(db.Token)
	authorizationInfoDto := dto.NewAuthorizationInfoDto(token.Allowed)

	if token.Allowed == db.TokenAllowedTypeALLOWED {
		locationReferencesDto, err := r.UnmarshalLocationReferencesDto(request.Body)

		if err != nil && err != io.EOF {
			render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
			return
		}

		// TODO: we should reject an authorization to an unknown location/evse/connector
		tokenAuthorization, err := r.CreateTokenAuthorization(ctx, *cred, token, locationReferencesDto)
		var displayText *coreDto.DisplayTextDto

		if err != nil {
			displayText = &coreDto.DisplayTextDto{
				Language: "en",
				Text: err.Error(),
			}
		}

		authorizationInfoDto = r.CreateAuthorizationInfoDto(ctx, token, tokenAuthorization, locationReferencesDto, displayText)

		if tokenAuthorization != nil && tokenAuthorization.Authorized {
			go r.waitForEvsesStatus(*cred, token, *tokenAuthorization, locationReferencesDto, db.EvseStatusCHARGING, 150)
		}
	}

	if err := render.Render(rw, request, transportation.OcpiSuccess(authorizationInfoDto)); err != nil {
		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}
