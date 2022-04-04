package token

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/displaytext"
	"github.com/satimoto/go-ocpi-api/internal/util"
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

func (r *TokenResolver) CreateAuthorizationInfoDto(ctx context.Context, token db.Token, tokenAuthorization *db.TokenAuthorization, location *LocationReferencesDto, info *displaytext.DisplayTextDto) *AuthorizationInfoDto {
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

func NewCreateTokenAuthorizationParams(tokenID int64) db.CreateTokenAuthorizationParams {
	return db.CreateTokenAuthorizationParams{
		TokenID:         tokenID,
		AuthorizationID: uuid.NewString(),
	}
}

type TokenDto struct {
	Uid          *string                `json:"uid"`
	Type         *db.TokenType          `json:"type"`
	AuthID       *string                `json:"auth_id"`
	VisualNumber *string                `json:"visual_number,omitempty"`
	Issuer       *string                `json:"issuer"`
	Valid        *bool                  `json:"valid"`
	Whitelist    *db.TokenWhitelistType `json:"whitelist"`
	Language     *string                `json:"language,omitempty"`
	LastUpdated  *time.Time             `json:"last_updated"`
}

func (r *TokenDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewTokenDto(token db.Token) *TokenDto {
	return &TokenDto{
		Uid:          &token.Uid,
		Type:         &token.Type,
		AuthID:       &token.AuthID,
		VisualNumber: util.NilString(token.VisualNumber.String),
		Issuer:       &token.Issuer,
		Valid:        &token.Valid,
		Whitelist:    &token.Whitelist,
		Language:     util.NilString(token.Language.String),
		LastUpdated:  &token.LastUpdated,
	}
}

func NewCreateTokenParams(dto *TokenDto) db.CreateTokenParams {
	return db.CreateTokenParams{
		Uid:          *dto.Uid,
		Type:         *dto.Type,
		AuthID:       *dto.AuthID,
		VisualNumber: util.SqlNullString(dto.VisualNumber),
		Issuer:       *dto.Issuer,
		Valid:        *dto.Valid,
		Whitelist:    *dto.Whitelist,
		Language:     util.SqlNullString(dto.Language),
		LastUpdated:  *dto.LastUpdated,
	}
}

func NewUpdateTokenByUidParams(token db.Token) db.UpdateTokenByUidParams {
	return db.UpdateTokenByUidParams{
		Uid:          token.Uid,
		Type:         token.Type,
		AuthID:       token.AuthID,
		VisualNumber: token.VisualNumber,
		Issuer:       token.Issuer,
		Allowed:      token.Allowed,
		Valid:        token.Valid,
		Whitelist:    token.Whitelist,
		Language:     token.Language,
		LastUpdated:  token.LastUpdated,
	}
}

func (r *TokenResolver) CreateTokenDto(ctx context.Context, token db.Token) *TokenDto {
	return NewTokenDto(token)
}

func (r *TokenResolver) CreateTokenListDto(ctx context.Context, tokens []db.Token) []render.Renderer {
	list := []render.Renderer{}
	for _, token := range tokens {
		list = append(list, r.CreateTokenDto(ctx, token))
	}
	return list
}
