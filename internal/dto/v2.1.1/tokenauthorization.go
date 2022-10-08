package dto

import (
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

type AuthorizationInfoDto struct {
	Allowed         *db.TokenAllowedType    `json:"allowed"`
	AuthorizationID *string                 `json:"authorization_id"`
	Location        *LocationReferencesDto  `json:"location,omitempty"`
	Info            *coreDto.DisplayTextDto `json:"info,omitempty"`
}

func (r *AuthorizationInfoDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewAuthorizationInfoDto(token db.Token) *AuthorizationInfoDto {
	return &AuthorizationInfoDto{
		Allowed: &token.Allowed,
	}
}

type LocationReferencesDto struct {
	LocationID   *string   `json:"location_id,omitempty"`
	EvseUids     []*string `json:"evse_uids,omitempty"`
	ConnectorIds []*string `json:"connector_ids,omitempty"`
}

func (r *LocationReferencesDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewLocationReferencesDto(locationID string) *LocationReferencesDto {
	return &LocationReferencesDto{
		LocationID: util.NilString(locationID),
	}
}
