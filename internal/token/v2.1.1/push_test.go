package token_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/nsf/jsondiff"
	"github.com/satimoto/go-datastore/pkg/db"
	dbMocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/middleware"
	notificationMocks "github.com/satimoto/go-ocpi/internal/notification/mocks"
	serviceMocks "github.com/satimoto/go-ocpi/internal/service/mocks"
	token "github.com/satimoto/go-ocpi/internal/token/v2.1.1"
	tokenMocks "github.com/satimoto/go-ocpi/internal/token/v2.1.1/mocks"
	transportationMocks "github.com/satimoto/go-ocpi/internal/transportation/mocks"
	"github.com/satimoto/go-ocpi/test/mocks"
)

func setupRoutes(tokenResolver *token.TokenResolver) *chi.Mux {
	router := chi.NewRouter()

	paginationContextRouter := router.With(middleware.Paginate)
	paginationContextRouter.Get("/", tokenResolver.ListTokens)

	router.Route("/{token_id}", func(tokenRouter chi.Router) {
		tokenContextRouter := tokenRouter.With(tokenResolver.TokenContext)
		tokenContextRouter.Post("/authorize", tokenResolver.AuthorizeToken)
	})

	return router
}

func TestTokenRequest(t *testing.T) {
	t.Run("Empty list tokens request", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		mockNotificationService := notificationMocks.NewService()
		mockOcpiService := transportationMocks.NewOcpiService(mockHTTPRequester)
		mockServices := serviceMocks.NewService(mockRepository, mockNotificationService, mockOcpiService)

		tokenResolver := tokenMocks.NewResolver(mockRepository, mockServices)
		tokenRoutes := setupRoutes(tokenResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodGet, "/", nil)

		if err != nil {
			t.Fatal("Creating 'GET /' request failed!")
		}

		tokenRoutes.ServeHTTP(responseRecorder, request)

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"data": [],
			"status_code": 1000,
			"status_message": "Success"
		}`), jsondiff.SupersetMatch)
	})

	t.Run("List tokens request", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		mockNotificationService := notificationMocks.NewService()
		mockOcpiService := transportationMocks.NewOcpiService(mockHTTPRequester)
		mockServices := serviceMocks.NewService(mockRepository, mockNotificationService, mockOcpiService)

		tokenResolver := tokenMocks.NewResolver(mockRepository, mockServices)
		tokenRoutes := setupRoutes(tokenResolver)
		responseRecorder := httptest.NewRecorder()

		tokens := []db.Token{}
		tokens = append(tokens, db.Token{
			Uid:          "TOKEN00001",
			Type:         db.TokenTypeOTHER,
			AuthID:       "DEBTCC30384929",
			VisualNumber: util.SqlNullString("DE-BTC-C30384929"),
			Issuer:       "Satimoto",
			Allowed:      db.TokenAllowedTypeALLOWED,
			Valid:        true,
			Whitelist:    db.TokenWhitelistTypeALLOWED,
			Language:     util.SqlNullString("en"),
			LastUpdated:  *util.ParseTime("2015-06-29T20:39:09Z", nil),
		})
		mockRepository.SetListTokensMockData(dbMocks.TokensMockData{Tokens: tokens, Error: nil})

		request, err := http.NewRequest(http.MethodGet, "/", nil)

		if err != nil {
			t.Fatal("Creating 'GET /' request failed!")
		}

		tokenRoutes.ServeHTTP(responseRecorder, request)

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"data": [{
				"uid": "TOKEN00001",
				"type": "OTHER",
				"auth_id": "DEBTCC30384929",
				"visual_number": "DE-BTC-C30384929",
				"issuer": "Satimoto",
				"valid": true,
				"whitelist": "ALLOWED",
				"language": "en",
				"last_updated": "2015-06-29T20:39:09Z"
			}],
			"status_code": 1000,
			"status_message": "Success"
		}`), jsondiff.SupersetMatch)
	})

	t.Run("Invalid route", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		mockNotificationService := notificationMocks.NewService()
		mockOcpiService := transportationMocks.NewOcpiService(mockHTTPRequester)
		mockServices := serviceMocks.NewService(mockRepository, mockNotificationService, mockOcpiService)

		tokenResolver := tokenMocks.NewResolver(mockRepository, mockServices)
		tokenRoutes := setupRoutes(tokenResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodGet, "/TOKEN00001", nil)

		if err != nil {
			t.Fatal("Creating 'GET /{token_id}' request failed!")
		}

		tokenRoutes.ServeHTTP(responseRecorder, request)

		if responseRecorder.Code != http.StatusNotFound {
			t.Fatal("Returned ", responseRecorder.Code, " instead of ", http.StatusNotFound)
		}
	})

	t.Run("Invalid authorize", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		mockNotificationService := notificationMocks.NewService()
		mockOcpiService := transportationMocks.NewOcpiService(mockHTTPRequester)
		mockServices := serviceMocks.NewService(mockRepository, mockNotificationService, mockOcpiService)

		tokenResolver := tokenMocks.NewResolver(mockRepository, mockServices)
		tokenRoutes := setupRoutes(tokenResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodPost, "/TOKEN00002/authorize", nil)

		if err != nil {
			t.Fatal("Creating 'GET /{token_id}/authorize' request failed!")
		}

		tokenRoutes.ServeHTTP(responseRecorder, request)

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"status_code": 2003,
			"status_message": "Unknown resource"
		}`), jsondiff.SupersetMatch)
	})

	t.Run("Authorize", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		mockNotificationService := notificationMocks.NewService()
		mockOcpiService := transportationMocks.NewOcpiService(mockHTTPRequester)
		mockServices := serviceMocks.NewService(mockRepository, mockNotificationService, mockOcpiService)

		tokenResolver := tokenMocks.NewResolver(mockRepository, mockServices)
		tokenRoutes := setupRoutes(tokenResolver)
		responseRecorder := httptest.NewRecorder()

		token := db.Token{
			Uid:          "TOKEN00001",
			Type:         db.TokenTypeOTHER,
			AuthID:       "DEBTCC30384929",
			VisualNumber: util.SqlNullString("DE-BTC-C30384929"),
			Issuer:       "Satimoto",
			Allowed:      db.TokenAllowedTypeALLOWED,
			Valid:        true,
			Whitelist:    db.TokenWhitelistTypeALLOWED,
			Language:     util.SqlNullString("en"),
			LastUpdated:  *util.ParseTime("2015-06-29T20:39:09Z", nil),
		}
		mockRepository.SetGetTokenByUidMockData(dbMocks.TokenMockData{Token: token, Error: nil})

		request, err := http.NewRequest(http.MethodPost, "/TOKEN00001/authorize", nil)

		if err != nil {
			t.Fatal("Creating 'GET /{token_id}/authorize' request failed!")
		}

		tokenRoutes.ServeHTTP(responseRecorder, request)

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"data": {
				"allowed": "ALLOWED"
			},
			"status_code": 1000,
			"status_message": "Success"
		}`), jsondiff.SupersetMatch)
	})

	t.Run("Authorize with location references", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		mockNotificationService := notificationMocks.NewService()
		mockOcpiService := transportationMocks.NewOcpiService(mockHTTPRequester)
		mockServices := serviceMocks.NewService(mockRepository, mockNotificationService, mockOcpiService)

		tokenResolver := tokenMocks.NewResolver(mockRepository, mockServices)
		tokenRoutes := setupRoutes(tokenResolver)
		responseRecorder := httptest.NewRecorder()

		token := db.Token{
			Uid:          "TOKEN00001",
			Type:         db.TokenTypeOTHER,
			AuthID:       "DEBTCC30384929",
			VisualNumber: util.SqlNullString("DE-BTC-C30384929"),
			Issuer:       "Satimoto",
			Allowed:      db.TokenAllowedTypeALLOWED,
			Valid:        true,
			Whitelist:    db.TokenWhitelistTypeALLOWED,
			Language:     util.SqlNullString("en"),
			LastUpdated:  *util.ParseTime("2015-06-29T20:39:09Z", nil),
		}
		mockRepository.SetGetTokenByUidMockData(dbMocks.TokenMockData{Token: token, Error: nil})

		request, err := http.NewRequest(http.MethodPost, "/TOKEN00001/authorize", bytes.NewBuffer([]byte(`{
			"location_id": "LOC0000001"
		}`)))
		request.Header.Set("Content-Type", "application/json")

		if err != nil {
			t.Fatal("Creating 'GET /{token_id}/authorize' request failed!")
		}

		tokenRoutes.ServeHTTP(responseRecorder, request)

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"data": {
				"allowed": "ALLOWED",
				"location": {
					"location_id": "LOC0000001"
				}
			},
			"status_code": 1000,
			"status_message": "Success"
		}`), jsondiff.SupersetMatch)
	})
	t.Run("Authorize with location references", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		mockNotificationService := notificationMocks.NewService()
		mockOcpiService := transportationMocks.NewOcpiService(mockHTTPRequester)
		mockServices := serviceMocks.NewService(mockRepository, mockNotificationService, mockOcpiService)

		tokenResolver := tokenMocks.NewResolver(mockRepository, mockServices)
		tokenRoutes := setupRoutes(tokenResolver)
		responseRecorder := httptest.NewRecorder()

		token := db.Token{
			Uid:          "TOKEN00001",
			Type:         db.TokenTypeOTHER,
			AuthID:       "DEBTCC30384929",
			VisualNumber: util.SqlNullString("DE-BTC-C30384929"),
			Issuer:       "Satimoto",
			Allowed:      db.TokenAllowedTypeALLOWED,
			Valid:        true,
			Whitelist:    db.TokenWhitelistTypeALLOWED,
			Language:     util.SqlNullString("en"),
			LastUpdated:  *util.ParseTime("2015-06-29T20:39:09Z", nil),
		}
		mockRepository.SetGetTokenByUidMockData(dbMocks.TokenMockData{Token: token, Error: nil})

		request, err := http.NewRequest(http.MethodPost, "/TOKEN00001/authorize", bytes.NewBuffer([]byte(`{
			"location_id": "LOC0000001"
		}`)))
		request.Header.Set("Content-Type", "application/json")

		if err != nil {
			t.Fatal("Creating 'GET /{token_id}/authorize' request failed!")
		}

		tokenRoutes.ServeHTTP(responseRecorder, request)

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"data": {
				"allowed": "ALLOWED",
				"location": {
					"location_id": "LOC0000001"
				}
			},
			"status_code": 1000,
			"status_message": "Success"
		}`), jsondiff.SupersetMatch)
	})

	t.Run("Authorize with location and evid references", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		mockNotificationService := notificationMocks.NewService()
		mockOcpiService := transportationMocks.NewOcpiService(mockHTTPRequester)
		mockServices := serviceMocks.NewService(mockRepository, mockNotificationService, mockOcpiService)

		tokenResolver := tokenMocks.NewResolver(mockRepository, mockServices)
		tokenRoutes := setupRoutes(tokenResolver)
		responseRecorder := httptest.NewRecorder()

		token := db.Token{
			Uid:          "TOKEN00001",
			Type:         db.TokenTypeOTHER,
			AuthID:       "DEBTCC30384929",
			VisualNumber: util.SqlNullString("DE-BTC-C30384929"),
			Issuer:       "Satimoto",
			Allowed:      db.TokenAllowedTypeALLOWED,
			Valid:        true,
			Whitelist:    db.TokenWhitelistTypeALLOWED,
			Language:     util.SqlNullString("en"),
			LastUpdated:  *util.ParseTime("2015-06-29T20:39:09Z", nil),
		}
		mockRepository.SetGetTokenByUidMockData(dbMocks.TokenMockData{Token: token, Error: nil})

		request, err := http.NewRequest(http.MethodPost, "/TOKEN00001/authorize", bytes.NewBuffer([]byte(`{
			"location_id": "LOC0000001",
			"evse_uids": ["EVSE000001", "EVSE000002"],
			"connector_ids": ["EVSE0000010001", "EVSE0000010002", "EVSE0000020001", "EVSE0000020002"]
		}`)))
		request.Header.Set("Content-Type", "application/json")

		if err != nil {
			t.Fatal("Creating 'GET /{token_id}/authorize' request failed!")
		}

		tokenRoutes.ServeHTTP(responseRecorder, request)

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"data": {
				"allowed": "ALLOWED",
				"location": {
					"location_id": "LOC0000001",
					"evse_uids": ["EVSE000001", "EVSE000002"],
					"connector_ids": ["EVSE0000010001", "EVSE0000010002", "EVSE0000020001", "EVSE0000020002"]
				}
			},
			"status_code": 1000,
			"status_message": "Success"
		}`), jsondiff.SupersetMatch)
	})
}
