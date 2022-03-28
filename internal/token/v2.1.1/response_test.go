package token_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"

	dbMocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-datastore/db"
	tokenMocks "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/util"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func TestCreateTokenPayload(t *testing.T) {
	ctx := context.Background()

	t.Run("Empty", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		tokenResolver := tokenMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

		loc := db.Token{}

		response := tokenResolver.CreateTokenPayload(ctx, loc)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"uid": "",
			"type": "",
			"auth_id": "",
			"issuer": "",
			"valid": false,
			"whitelist": "",
			"last_updated": "0001-01-01T00:00:00Z"
		}`))
	})

	t.Run("Basic token", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		tokenResolver := tokenMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

		tok := db.Token{
			Uid:          "TOKEN00001",
			Type:         db.TokenTypeRFID,
			AuthID:       "DEBTCC30384929",
			Issuer:       "Satimoto",
			Allowed:      db.TokenAllowedTypeEXPIRED,
			Valid:        false,
			Whitelist:    db.TokenWhitelistTypeNEVER,
			LastUpdated:  *util.ParseTime("2015-06-29T20:39:09Z"),
		}

		response := tokenResolver.CreateTokenPayload(ctx, tok)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"uid": "TOKEN00001",
			"type": "RFID",
			"auth_id": "DEBTCC30384929",
			"issuer": "Satimoto",
			"valid": false,
			"whitelist": "NEVER",
			"last_updated": "2015-06-29T20:39:09Z"
		}`))
	})

	t.Run("With visual number and language", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		tokenResolver := tokenMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

		tok := db.Token{
			Uid:          "TOKEN00001",
			Type:         db.TokenTypeOTHER,
			AuthID:       "DEBTCC30384929",
			VisualNumber: sql.NullString{String: "DE-BTC-C30384929"},
			Issuer:       "Satimoto",
			Allowed:      db.TokenAllowedTypeNOCREDIT,
			Valid:        true,
			Whitelist:    db.TokenWhitelistTypeALLOWEDOFFLINE,
			LastUpdated:  *util.ParseTime("2015-06-29T20:39:09Z"),
		}

		response := tokenResolver.CreateTokenPayload(ctx, tok)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"uid": "TOKEN00001",
			"type": "OTHER",
			"auth_id": "DEBTCC30384929",
			"visual_number": "DE-BTC-C30384929",
			"issuer": "Satimoto",
			"valid": true,
			"whitelist": "ALLOWED_OFFLINE",
			"last_updated": "2015-06-29T20:39:09Z"
		}`))
	})

	t.Run("With visual number and language", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		tokenResolver := tokenMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

		tok := db.Token{
			Uid:          "TOKEN00001",
			Type:         db.TokenTypeOTHER,
			AuthID:       "DEBTCC30384929",
			VisualNumber: sql.NullString{String: "DE-BTC-C30384929"},
			Issuer:       "Satimoto",
			Allowed:      db.TokenAllowedTypeALLOWED,
			Valid:        true,
			Whitelist:    db.TokenWhitelistTypeALWAYS,
			Language:     sql.NullString{String: "en"},
			LastUpdated:  *util.ParseTime("2015-06-29T20:39:09Z"),
		}

		response := tokenResolver.CreateTokenPayload(ctx, tok)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"uid": "TOKEN00001",
			"type": "OTHER",
			"auth_id": "DEBTCC30384929",
			"visual_number": "DE-BTC-C30384929",
			"issuer": "Satimoto",
			"valid": true,
			"whitelist": "ALWAYS",
			"language": "en",
			"last_updated": "2015-06-29T20:39:09Z"
		}`))
	})
}
