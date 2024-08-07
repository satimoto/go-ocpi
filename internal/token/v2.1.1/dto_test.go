package token_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/satimoto/go-datastore/pkg/db"
	dbMocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/util"
	tokenMocks "github.com/satimoto/go-ocpi/internal/token/v2.1.1/mocks"
	transportationMocks "github.com/satimoto/go-ocpi/internal/transportation/mocks"
	"github.com/satimoto/go-ocpi/test/mocks"
	notificationMocks "github.com/satimoto/go-ocpi/internal/notification/mocks"
	serviceMocks "github.com/satimoto/go-ocpi/internal/service/mocks"
)

func TestCreateTokenDto(t *testing.T) {
	ctx := context.Background()

	t.Run("Empty", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		mockNotificationService := notificationMocks.NewService()
		mockOcpiService := transportationMocks.NewOcpiService(mockHTTPRequester)
		mockServices := serviceMocks.NewService(mockRepository, mockNotificationService, mockOcpiService)

		tokenResolver := tokenMocks.NewResolver(mockRepository, mockServices)

		loc := db.Token{}

		response := tokenResolver.CreateTokenDto(ctx, loc)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"uid": "",
			"type": "",
			"auth_id": "",
			"issuer": "",
			"valid": false,
			"whitelist": "",
			"last_updated": null
		}`))
	})

	t.Run("Basic token", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		mockNotificationService := notificationMocks.NewService()
		mockOcpiService := transportationMocks.NewOcpiService(mockHTTPRequester)
		mockServices := serviceMocks.NewService(mockRepository, mockNotificationService, mockOcpiService)

		tokenResolver := tokenMocks.NewResolver(mockRepository, mockServices)

		tok := db.Token{
			Uid:         "TOKEN00001",
			Type:        db.TokenTypeRFID,
			AuthID:      "DEBTCC30384929",
			Issuer:      "Satimoto",
			Allowed:     db.TokenAllowedTypeEXPIRED,
			Valid:       false,
			Whitelist:   db.TokenWhitelistTypeNEVER,
			LastUpdated: *util.ParseTime("2015-06-29T20:39:09Z", nil),
		}

		response := tokenResolver.CreateTokenDto(ctx, tok)
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
		mockNotificationService := notificationMocks.NewService()
		mockOcpiService := transportationMocks.NewOcpiService(mockHTTPRequester)
		mockServices := serviceMocks.NewService(mockRepository, mockNotificationService, mockOcpiService)

		tokenResolver := tokenMocks.NewResolver(mockRepository, mockServices)

		tok := db.Token{
			Uid:          "TOKEN00001",
			Type:         db.TokenTypeOTHER,
			AuthID:       "DEBTCC30384929",
			VisualNumber: util.SqlNullString("DE-BTC-C30384929"),
			Issuer:       "Satimoto",
			Allowed:      db.TokenAllowedTypeNOCREDIT,
			Valid:        true,
			Whitelist:    db.TokenWhitelistTypeALLOWEDOFFLINE,
			LastUpdated:  *util.ParseTime("2015-06-29T20:39:09Z", nil),
		}

		response := tokenResolver.CreateTokenDto(ctx, tok)
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
		mockNotificationService := notificationMocks.NewService()
		mockOcpiService := transportationMocks.NewOcpiService(mockHTTPRequester)
		mockServices := serviceMocks.NewService(mockRepository, mockNotificationService, mockOcpiService)

		tokenResolver := tokenMocks.NewResolver(mockRepository, mockServices)

		tok := db.Token{
			Uid:          "TOKEN00001",
			Type:         db.TokenTypeOTHER,
			AuthID:       "DEBTCC30384929",
			VisualNumber: util.SqlNullString("DE-BTC-C30384929"),
			Issuer:       "Satimoto",
			Allowed:      db.TokenAllowedTypeALLOWED,
			Valid:        true,
			Whitelist:    db.TokenWhitelistTypeALWAYS,
			Language:     util.SqlNullString("en"),
			LastUpdated:  *util.ParseTime("2015-06-29T20:39:09Z", nil),
		}

		response := tokenResolver.CreateTokenDto(ctx, tok)
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
