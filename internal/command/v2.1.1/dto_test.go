package command_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	dbMocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-datastore/db"
	commandMocks "github.com/satimoto/go-ocpi-api/internal/command/v2.1.1/mocks"
	ocpiMocks "github.com/satimoto/go-ocpi-api/internal/ocpi/mocks"
	"github.com/satimoto/go-ocpi-api/internal/util"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

var (
	apiDomain = "http://localhost:9001"
)

func TestCreateCommandReservationDto(t *testing.T) {
	ctx := context.Background()
	os.Setenv("API_DOMAIN", apiDomain)
	defer os.Unsetenv("API_DOMAIN")

	t.Run("Empty", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolver(mockRepository, ocpiMocks.NewOCPIRequester(mockHTTPRequester))

		tok := db.Token{}
		mockRepository.SetGetTokenMockData(dbMocks.TokenMockData{Token: tok, Error: nil})

		cr := db.CommandReservation{
			TokenID: 1,
		}

		response := commandResolver.CreateCommandReservationDto(ctx, cr)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"response_url": "http://localhost:9001/2.1.1/commands/RESERVE_NOW/0",
			"token": {
				"auth_id": "",
				"issuer": "",
				"last_updated": "0001-01-01T00:00:00Z",
				"type": "",
				"uid": "",
				"valid": false,
				"whitelist": ""
			},
			"expiry_date": "0001-01-01T00:00:00Z",
			"reservation_id": 0,
			"location_id": ""
		}`))
	})

	t.Run("Command reservation", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolver(mockRepository, ocpiMocks.NewOCPIRequester(mockHTTPRequester))

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
			LastUpdated:  *util.ParseTime("2015-06-29T20:39:09Z"),
		}
		mockRepository.SetGetTokenMockData(dbMocks.TokenMockData{Token: tok, Error: nil})

		cr := db.CommandReservation{
			ID:            1,
			Status:        db.CommandResponseTypeREQUESTED,
			TokenID:       1,
			ExpiryDate:    *util.ParseTime("2015-06-29T20:39:09Z"),
			ReservationID: 2,
			LocationID:    "LOC00001",
		}

		response := commandResolver.CreateCommandReservationDto(ctx, cr)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"response_url": "http://localhost:9001/2.1.1/commands/RESERVE_NOW/1",
			"token": {
				"uid": "TOKEN00001",
				"type": "OTHER",
				"auth_id": "DEBTCC30384929",
				"visual_number": "DE-BTC-C30384929",
				"issuer": "Satimoto",
				"valid": true,
				"whitelist": "ALWAYS",
				"language": "en",
				"last_updated": "2015-06-29T20:39:09Z"
			},
			"expiry_date": "2015-06-29T20:39:09Z",
			"reservation_id": 2,
			"location_id": "LOC00001"
		}`))
	})
}

func TestCreateCommandStartDto(t *testing.T) {
	ctx := context.Background()
	os.Setenv("API_DOMAIN", apiDomain)
	defer os.Unsetenv("API_DOMAIN")

	t.Run("Empty", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolver(mockRepository, ocpiMocks.NewOCPIRequester(mockHTTPRequester))

		tok := db.Token{}
		mockRepository.SetGetTokenMockData(dbMocks.TokenMockData{Token: tok, Error: nil})

		cr := db.CommandStart{
			TokenID: 1,
		}

		response := commandResolver.CreateCommandStartDto(ctx, cr)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"response_url": "http://localhost:9001/2.1.1/commands/START_SESSION/0",
			"token": {
				"auth_id": "",
				"issuer": "",
				"last_updated": "0001-01-01T00:00:00Z",
				"type": "",
				"uid": "",
				"valid": false,
				"whitelist": ""
			},
			"location_id": ""
		}`))
	})

	t.Run("Command start", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolver(mockRepository, ocpiMocks.NewOCPIRequester(mockHTTPRequester))

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
			LastUpdated:  *util.ParseTime("2015-06-29T20:39:09Z"),
		}
		mockRepository.SetGetTokenMockData(dbMocks.TokenMockData{Token: tok, Error: nil})

		cr := db.CommandStart{
			ID:         1,
			Status:     db.CommandResponseTypeREQUESTED,
			TokenID:    1,
			LocationID: "LOC00001",
		}

		response := commandResolver.CreateCommandStartDto(ctx, cr)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"response_url": "http://localhost:9001/2.1.1/commands/START_SESSION/1",
			"token": {
				"uid": "TOKEN00001",
				"type": "OTHER",
				"auth_id": "DEBTCC30384929",
				"visual_number": "DE-BTC-C30384929",
				"issuer": "Satimoto",
				"valid": true,
				"whitelist": "ALWAYS",
				"language": "en",
				"last_updated": "2015-06-29T20:39:09Z"
			},
			"location_id": "LOC00001"
		}`))
	})
}

func TestCreateCommandStopDto(t *testing.T) {
	ctx := context.Background()
	os.Setenv("API_DOMAIN", apiDomain)
	defer os.Unsetenv("API_DOMAIN")

	t.Run("Empty", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolver(mockRepository, ocpiMocks.NewOCPIRequester(mockHTTPRequester))

		cr := db.CommandStop{}

		response := commandResolver.CreateCommandStopDto(ctx, cr)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"response_url": "http://localhost:9001/2.1.1/commands/STOP_SESSION/0",
			"session_id": ""
		}`))
	})

	t.Run("Command stop", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolver(mockRepository, ocpiMocks.NewOCPIRequester(mockHTTPRequester))

		cr := db.CommandStop{
			ID:        1,
			Status:    db.CommandResponseTypeREQUESTED,
			SessionID: "SESSION001",
		}

		response := commandResolver.CreateCommandStopDto(ctx, cr)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"response_url": "http://localhost:9001/2.1.1/commands/STOP_SESSION/1",
			"session_id": "SESSION001"
		}`))
	})
}

func TestCreateCommandUnlockDto(t *testing.T) {
	ctx := context.Background()
	os.Setenv("API_DOMAIN", apiDomain)
	defer os.Unsetenv("API_DOMAIN")

	t.Run("Empty", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolver(mockRepository, ocpiMocks.NewOCPIRequester(mockHTTPRequester))

		cr := db.CommandUnlock{}

		response := commandResolver.CreateCommandUnlockDto(ctx, cr)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"response_url": "http://localhost:9001/2.1.1/commands/UNLOCK_CONNECTOR/0",
			"location_id": "",
			"evse_uid": "",
			"connector_id": ""
		}`))
	})

	t.Run("Command unlock", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolver(mockRepository, ocpiMocks.NewOCPIRequester(mockHTTPRequester))

		cr := db.CommandUnlock{
			ID:          1,
			Status:      db.CommandResponseTypeREQUESTED,
			LocationID:  "LOC00001",
			EvseUid:     "EVSE00001",
			ConnectorID: "CONN00001",
		}

		response := commandResolver.CreateCommandUnlockDto(ctx, cr)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"response_url": "http://localhost:9001/2.1.1/commands/UNLOCK_CONNECTOR/1",
			"location_id": "LOC00001",
			"evse_uid": "EVSE00001",
			"connector_id": "CONN00001"
		}`))
	})
}
