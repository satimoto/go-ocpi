package session_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/nsf/jsondiff"
	"github.com/satimoto/go-datastore/pkg/db"
	dbMocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/util"
	notificationMocks "github.com/satimoto/go-ocpi/internal/notification/mocks"
	serviceMocks "github.com/satimoto/go-ocpi/internal/service/mocks"
	session "github.com/satimoto/go-ocpi/internal/session/v2.1.1"
	sessionMocks "github.com/satimoto/go-ocpi/internal/session/v2.1.1/mocks"
	transportationMocks "github.com/satimoto/go-ocpi/internal/transportation/mocks"
	"github.com/satimoto/go-ocpi/test/mocks"
)

func setupRoutes(sessionResolver *session.SessionResolver) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/{country_code}/{party_id}/{session_id}", func(sessionRouter chi.Router) {
		sessionRouter.Put("/", sessionResolver.UpdateSession)

		sessionContextRouter := sessionRouter.With(sessionResolver.SessionContext)
		sessionContextRouter.Get("/", sessionResolver.GetSession)
		sessionContextRouter.Patch("/", sessionResolver.UpdateSession)
	})

	return router
}

func TestSessionRequest(t *testing.T) {
	t.Run("Invalid route", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		mockNotificationService := notificationMocks.NewService()
		mockOcpiService := transportationMocks.NewOcpiService(mockHTTPRequester)
		mockServices := serviceMocks.NewService(mockRepository, mockNotificationService, mockOcpiService)

		sessionResolver := sessionMocks.NewResolver(mockRepository, mockServices)
		sessionRoutes := setupRoutes(sessionResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodGet, "/DE/ABC", nil)

		if err != nil {
			t.Fatal("Creating 'GET /{country_code}/{party_id}' request failed!")
		}

		sessionRoutes.ServeHTTP(responseRecorder, request)

		if responseRecorder.Code != http.StatusNotFound {
			t.Fatal("Returned ", responseRecorder.Code, " instead of ", http.StatusNotFound)
		}
	})

	t.Run("Get pending session", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		mockNotificationService := notificationMocks.NewService()
		mockOcpiService := transportationMocks.NewOcpiService(mockHTTPRequester)
		mockServices := serviceMocks.NewService(mockRepository, mockNotificationService, mockOcpiService)

		sessionResolver := sessionMocks.NewResolver(mockRepository, mockServices)
		sessionRoutes := setupRoutes(sessionResolver)
		responseRecorder := httptest.NewRecorder()

		sess := db.Session{
			Uid:           "SESSION0001",
			StartDatetime: *util.ParseTime("2015-06-29T22:39:09Z", nil),
			Kwh:           0,
			AuthID:        "DE8ACC12E46L89",
			AuthMethod:    db.AuthMethodTypeAUTHREQUEST,
			Currency:      "EUR",
			Status:        db.SessionStatusTypePENDING,
			LastUpdated:   *util.ParseTime("2015-06-29T22:39:09Z", nil),
		}
		mockRepository.SetGetSessionByUidMockData(dbMocks.SessionMockData{Session: sess, Error: nil})

		request, err := http.NewRequest(http.MethodGet, "/DE/ABC/SESSION0001", nil)

		if err != nil {
			t.Fatal("Creating 'GET /{country_code}/{party_id}/{session_id}' request failed!")
		}

		sessionRoutes.ServeHTTP(responseRecorder, request)

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"data": {
				"id": "SESSION0001",
				"start_datetime": "2015-06-29T22:39:09Z",
				"kwh": 0,
				"auth_id": "DE8ACC12E46L89",
				"auth_method": "AUTH_REQUEST",
				"location": null,
				"currency": "EUR",
				"charging_periods": [],
				"status": "PENDING",
				"last_updated": "2015-06-29T22:39:09Z"
			},
			"status_code": 1000,
			"status_message": "Success"
		}`), jsondiff.SupersetMatch)
	})

	t.Run("Get active session", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		mockNotificationService := notificationMocks.NewService()
		mockOcpiService := transportationMocks.NewOcpiService(mockHTTPRequester)
		mockServices := serviceMocks.NewService(mockRepository, mockNotificationService, mockOcpiService)

		sessionResolver := sessionMocks.NewResolver(mockRepository, mockServices)
		sessionRoutes := setupRoutes(sessionResolver)
		responseRecorder := httptest.NewRecorder()

		chargingPeriods := []db.ChargingPeriod{}
		chargingPeriods = append(chargingPeriods, db.ChargingPeriod{
			StartDateTime: *util.ParseTime("2015-06-29T22:39:09Z", nil),
		})
		mockRepository.SetListSessionChargingPeriodsMockData(dbMocks.ChargingPeriodsMockData{ChargingPeriods: chargingPeriods, Error: nil})

		dimensions := []db.ChargingPeriodDimension{}
		dimensions = append(dimensions, db.ChargingPeriodDimension{
			Type:   db.ChargingPeriodDimensionTypeFLAT,
			Volume: 1,
		})
		dimensions = append(dimensions, db.ChargingPeriodDimension{
			Type:   db.ChargingPeriodDimensionTypeTIME,
			Volume: 1.973,
		})
		mockRepository.SetListChargingPeriodDimensionsMockData(dbMocks.ChargingPeriodDimensionsMockData{ChargingPeriodDimensions: dimensions, Error: nil})

		sess := db.Session{
			Uid:           "SESSION0001",
			StartDatetime: *util.ParseTime("2015-06-29T22:39:09Z", nil),
			Kwh:           15.342,
			AuthID:        "DE8ACC12E46L89",
			AuthMethod:    db.AuthMethodTypeAUTHREQUEST,
			Currency:      "EUR",
			Status:        db.SessionStatusTypeACTIVE,
			LastUpdated:   *util.ParseTime("2015-06-29T22:39:09Z", nil),
		}
		mockRepository.SetGetSessionByUidMockData(dbMocks.SessionMockData{Session: sess, Error: nil})

		request, err := http.NewRequest(http.MethodGet, "/DE/ABC/SESSION0001", nil)

		if err != nil {
			t.Fatal("Creating 'GET /{country_code}/{party_id}/{session_id}' request failed!")
		}

		sessionRoutes.ServeHTTP(responseRecorder, request)

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"data": {
				"id": "SESSION0001",
				"start_datetime": "2015-06-29T22:39:09Z",
				"kwh": 15.342,
				"auth_id": "DE8ACC12E46L89",
				"auth_method": "AUTH_REQUEST",
				"location": null,
				"currency": "EUR",
				"charging_periods": [{
					"start_date_time": "2015-06-29T22:39:09Z",
					"dimensions": [{
						"type": "FLAT",
						"volume": 1
					}, {
						"type": "TIME",
						"volume": 1.973
					}]
				}],
				"status": "ACTIVE",
				"last_updated": "2015-06-29T22:39:09Z"
			},
			"status_code": 1000,
			"status_message": "Success"
		}`), jsondiff.SupersetMatch)
	})

	t.Run("Patch pending session", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		mockNotificationService := notificationMocks.NewService()
		mockOcpiService := transportationMocks.NewOcpiService(mockHTTPRequester)
		mockServices := serviceMocks.NewService(mockRepository, mockNotificationService, mockOcpiService)

		sessionResolver := sessionMocks.NewResolver(mockRepository, mockServices)
		sessionRoutes := setupRoutes(sessionResolver)
		responseRecorder := httptest.NewRecorder()

		chargingPeriods := []db.ChargingPeriod{}
		chargingPeriods = append(chargingPeriods, db.ChargingPeriod{
			StartDateTime: *util.ParseTime("2015-06-29T22:39:09Z", nil),
		})
		mockRepository.SetListSessionChargingPeriodsMockData(dbMocks.ChargingPeriodsMockData{ChargingPeriods: chargingPeriods, Error: nil})

		dimensions := []db.ChargingPeriodDimension{}
		dimensions = append(dimensions, db.ChargingPeriodDimension{
			Type:   db.ChargingPeriodDimensionTypeFLAT,
			Volume: 1,
		})
		dimensions = append(dimensions, db.ChargingPeriodDimension{
			Type:   db.ChargingPeriodDimensionTypeTIME,
			Volume: 1.973,
		})
		mockRepository.SetListChargingPeriodDimensionsMockData(dbMocks.ChargingPeriodDimensionsMockData{ChargingPeriodDimensions: dimensions, Error: nil})

		sess := db.Session{
			Uid:           "SESSION0001",
			StartDatetime: *util.ParseTime("2015-06-29T22:39:09Z", nil),
			Kwh:           15.342,
			AuthID:        "DE8ACC12E46L89",
			AuthMethod:    db.AuthMethodTypeAUTHREQUEST,
			LocationID:    1,
			Currency:      "EUR",
			Status:        db.SessionStatusTypeACTIVE,
			LastUpdated:   *util.ParseTime("2015-06-29T22:39:09Z", nil),
		}
		// Push context
		mockRepository.SetGetSessionByUidMockData(dbMocks.SessionMockData{Session: sess, Error: nil})
		// Push replace
		mockRepository.SetGetSessionByUidMockData(dbMocks.SessionMockData{Session: sess, Error: nil})

		request, err := http.NewRequest(http.MethodPatch, "/DE/ABC/SESSION0001", bytes.NewBuffer([]byte(`{
			"kwh": 16.1
		}`)))

		if err != nil {
			t.Fatal("Creating 'PATCH /{country_code}/{party_id}/{session_id}' request failed!")
		}

		requestCtx := request.Context()
		requestCtx = context.WithValue(requestCtx, "credential", &db.Credential{ID: 1, CountryCode: "FR", PartyID: "GER"})
		sessionRoutes.ServeHTTP(responseRecorder, request.WithContext(requestCtx))

		sessionParams, err := mockRepository.GetUpdateSessionByUidMockData()
		paramsJson, _ := json.Marshal(sessionParams)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"uid": "SESSION0001",
			"authorizationID": {"String": "", "Valid": false},
			"startDatetime": "2015-06-29T22:39:09Z",
			"endDatetime": {"Time": "0001-01-01T00:00:00Z", "Valid": false},
			"invoiceRequestID": {"Int64": 0, "Valid": false},
			"kwh": 16.1,
			"authMethod": "AUTH_REQUEST",
			"meterID": {"String": "", "Valid": false},
			"currency": "EUR",
			"totalCost": {"Float64": 0, "Valid": false},
			"status": "ACTIVE",
			"lastUpdated": "2015-06-29T22:39:09Z"
		}`))

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"status_code": 1000,
			"status_message": "Success"
		}`), jsondiff.SupersetMatch)
	})
}
