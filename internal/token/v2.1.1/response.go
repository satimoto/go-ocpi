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

type AuthorizationInfoPayload struct {
	Allowed         *db.TokenAllowedType            `json:"allowed"`
	AuthorizationID *string                         `json:"authorization_id"`
	Location        *LocationReferencesPayload      `json:"location,omitempty"`
	Info            *displaytext.DisplayTextPayload `json:"info,omitempty"`
}

func (r *AuthorizationInfoPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewAuthorizationInfoPayload(token db.Token) *AuthorizationInfoPayload {
	return &AuthorizationInfoPayload{
		Allowed: &token.Allowed,
	}
}

func (r *TokenResolver) CreateAuthorizationInfoPayload(ctx context.Context, token db.Token, tokenAuthorization *db.TokenAuthorization, location *LocationReferencesPayload, info *displaytext.DisplayTextPayload) *AuthorizationInfoPayload {
	response := NewAuthorizationInfoPayload(token)

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

type LocationReferencesPayload struct {
	LocationID   *string   `json:"location_id"`
	EvseUids     []*string `json:"evse_uids,omitempty"`
	ConnectorIds []*string `json:"connector_ids,omitempty"`
}

func (r *LocationReferencesPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCreateTokenAuthorizationParams(tokenID int64) db.CreateTokenAuthorizationParams {
	return db.CreateTokenAuthorizationParams{
		TokenID:         tokenID,
		AuthorizationID: uuid.NewString(),
	}
}

type TokenPayload struct {
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

func (r *TokenPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewTokenPayload(token db.Token) *TokenPayload {
	return &TokenPayload{
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

func NewCreateTokenParams(payload *TokenPayload) db.CreateTokenParams {
	return db.CreateTokenParams{
		Uid:          *payload.Uid,
		Type:         *payload.Type,
		AuthID:       *payload.AuthID,
		VisualNumber: util.SqlNullString(payload.VisualNumber),
		Issuer:       *payload.Issuer,
		Valid:        *payload.Valid,
		Whitelist:    *payload.Whitelist,
		Language:     util.SqlNullString(payload.Language),
		LastUpdated:  *payload.LastUpdated,
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

func (r *TokenResolver) CreateTokenPayload(ctx context.Context, token db.Token) *TokenPayload {
	return NewTokenPayload(token)
}

func (r *TokenResolver) CreateTokenListPayload(ctx context.Context, tokens []db.Token) []render.Renderer {
	list := []render.Renderer{}
	for _, token := range tokens {
		list = append(list, r.CreateTokenPayload(ctx, token))
	}
	return list
}
