package tokenauthorization

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi-api/internal/displaytext"
)

type AuthorizationInfoDto struct {
	Allowed         *db.TokenAllowedType        `json:"allowed"`
	AuthorizationID *string                     `json:"authorization_id"`
	Location        *LocationReferencesDto      `json:"location,omitempty"`
	Info            *displaytext.DisplayTextDto `json:"info,omitempty"`
}

func (r *AuthorizationInfoDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewAuthorizationInfoDto(token db.Token) *AuthorizationInfoDto {
	return &AuthorizationInfoDto{
		Allowed: &token.Allowed,
	}
}

func (r *TokenAuthorizationResolver) CreateAuthorizationInfoDto(ctx context.Context, token db.Token, tokenAuthorization *db.TokenAuthorization, location *LocationReferencesDto, info *displaytext.DisplayTextDto) *AuthorizationInfoDto {
	response := NewAuthorizationInfoDto(token)

	if tokenAuthorization != nil {
		response.AuthorizationID = &tokenAuthorization.AuthorizationID
	}

	if location != nil {
		response.Location = location
	}

	if info != nil {
		response.Info = info
	}

	return response
}

type LocationReferencesDto struct {
	LocationID   *string   `json:"location_id"`
	EvseUids     []*string `json:"evse_uids,omitempty"`
	ConnectorIds []*string `json:"connector_ids,omitempty"`
}

func (r *LocationReferencesDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
