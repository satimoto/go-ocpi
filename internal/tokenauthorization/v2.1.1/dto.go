package tokenauthorization

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)

func (r *TokenAuthorizationResolver) CreateAuthorizationInfoDto(ctx context.Context, token db.Token, tokenAuthorization *db.TokenAuthorization, location *dto.LocationReferencesDto, info *coreDto.DisplayTextDto) *dto.AuthorizationInfoDto {
	if tokenAuthorization != nil {
		response := dto.NewAuthorizationInfoDto(token.Allowed)
		response.AuthorizationID = &tokenAuthorization.AuthorizationID
	
		if location != nil {
			response.Location = location
		}
	
		if info != nil {
			response.Info = info
		}
	
		return response
	}

	return dto.NewAuthorizationInfoDto(db.TokenAllowedTypeNOTALLOWED)
}
