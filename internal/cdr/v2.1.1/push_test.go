package cdr_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi"
	"github.com/nsf/jsondiff"
	dbMocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-datastore/db"
	cdr "github.com/satimoto/go-ocpi-api/internal/cdr/v2.1.1"
	cdrMocks "github.com/satimoto/go-ocpi-api/internal/cdr/v2.1.1/mocks"
	ocpiMocks "github.com/satimoto/go-ocpi-api/internal/ocpi/mocks"
	"github.com/satimoto/go-ocpi-api/internal/util"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

var (
	apiDomain = "http://localhost:9001"
)

func setupRoutes(cdrResolver *cdr.CdrResolver) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/", func(credentialRouter chi.Router) {
		credentialRouter.Post("/", cdrResolver.PostCdr)

		credentialRouter.Route("/{cdr_id}", func(cdrRouter chi.Router) {
			cdrContextRouter := cdrRouter.With(cdrResolver.CdrContext)
			cdrContextRouter.Get("/", cdrResolver.GetCdr)
		})
	})

	return router
}

func TestCdrRequest(t *testing.T) {
	os.Setenv("API_DOMAIN", apiDomain)
	defer os.Unsetenv("API_DOMAIN")

	t.Run("Invalid route", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		cdrResolver := cdrMocks.NewResolver(mockRepository, ocpiMocks.NewOCPIRequester(mockHTTPRequester))
		cdrRoutes := setupRoutes(cdrResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest("GET", "/", nil)

		if err != nil {
			t.Fatal("Creating 'GET /' request failed!")
		}

		cdrRoutes.ServeHTTP(responseRecorder, request)

		if responseRecorder.Code != http.StatusMethodNotAllowed {
			t.Fatal("Returned ", responseRecorder.Code, " instead of ", http.StatusMethodNotAllowed)
		}
	})

	t.Run("Get cdr 1", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		cdrResolver := cdrMocks.NewResolver(mockRepository, ocpiMocks.NewOCPIRequester(mockHTTPRequester))
		cdrRoutes := setupRoutes(cdrResolver)
		responseRecorder := httptest.NewRecorder()

		c := db.Cdr{
			ID:            1,
			Uid:           "CDR0001",
			StartDateTime: *util.ParseTime("2015-06-29T21:39:09Z"),
			StopDateTime:  util.SqlNullTime(util.ParseTime("2015-06-29T21:39:09Z")),
			AuthID:        "DE8ACC12E46L89",
			AuthMethod:    db.AuthMethodTypeAUTHREQUEST,
			Currency:      "EUR",
			TotalCost:     4.0,
			TotalEnergy:   15.342,
			TotalTime:     1.973,
			LastUpdated:   *util.ParseTime("2015-06-29T22:01:13Z"),
		}
		mockRepository.SetGetCdrByUidMockData(dbMocks.CdrMockData{Cdr: c, Error: nil})

		request, err := http.NewRequest("GET", "/CDR0001", nil)

		if err != nil {
			t.Fatal("Creating 'GET /{cdr_id}' request failed!")
		}

		cdrRoutes.ServeHTTP(responseRecorder, request)

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"data": {
				"id": "CDR0001",
				"start_date_time": "2015-06-29T21:39:09Z",
				"stop_date_time": "2015-06-29T21:39:09Z",
				"auth_id": "DE8ACC12E46L89",
				"auth_method": "AUTH_REQUEST",
				"location": null,
				"currency": "EUR",
				"tariffs": [],
				"charging_periods": [],
				"total_cost": 4,
				"total_energy": 15.342,
				"total_time": 1.973,
				"last_updated": "2015-06-29T22:01:13Z"
			},
			"status_code": 1000,
			"status_message": "Success"
		}`), jsondiff.SupersetMatch)
	})

	t.Run("Get cdr 2", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		cdrResolver := cdrMocks.NewResolver(mockRepository, ocpiMocks.NewOCPIRequester(mockHTTPRequester))
		cdrRoutes := setupRoutes(cdrResolver)
		responseRecorder := httptest.NewRecorder()

		chargingPeriods := []db.ChargingPeriod{}
		chargingPeriods = append(chargingPeriods, db.ChargingPeriod{
			StartDateTime: *util.ParseTime("2015-06-29T22:39:09Z"),
		})
		mockRepository.SetListCdrChargingPeriodsMockData(dbMocks.ChargingPeriodsMockData{ChargingPeriods: chargingPeriods, Error: nil})

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

		c := db.Cdr{
			ID:               2,
			Uid:              "CDR0002",
			StartDateTime:    *util.ParseTime("2015-06-29T21:39:09Z"),
			StopDateTime:     util.SqlNullTime(util.ParseTime("2015-06-29T21:39:09Z")),
			AuthID:           "DE8ACC12E46L89",
			AuthMethod:       db.AuthMethodTypeAUTHREQUEST,
			Currency:         "EUR",
			TotalCost:        4.0,
			TotalEnergy:      15.342,
			TotalTime:        1.973,
			TotalParkingTime: util.SqlNullFloat64(45),
			LastUpdated:      *util.ParseTime("2015-06-29T22:01:13Z"),
		}
		mockRepository.SetGetCdrByUidMockData(dbMocks.CdrMockData{Cdr: c, Error: nil})

		request, err := http.NewRequest("GET", "/CDR0002", nil)

		if err != nil {
			t.Fatal("Creating 'GET /{cdr_id}' request failed!")
		}

		cdrRoutes.ServeHTTP(responseRecorder, request)

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"data": {
				"id": "CDR0002",
				"start_date_time": "2015-06-29T21:39:09Z",
				"stop_date_time": "2015-06-29T21:39:09Z",
				"auth_id": "DE8ACC12E46L89",
				"auth_method": "AUTH_REQUEST",
				"location": null,
				"currency": "EUR",
				"tariffs": [],
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
				"total_cost": 4,
				"total_energy": 15.342,
				"total_time": 1.973,
				"total_parking_time": 45,
				"last_updated": "2015-06-29T22:01:13Z"
			},
			"status_code": 1000,
			"status_message": "Success"
		}`), jsondiff.SupersetMatch)
	})

	t.Run("Post cdr", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		cdrResolver := cdrMocks.NewResolver(mockRepository, ocpiMocks.NewOCPIRequester(mockHTTPRequester))
		cdrRoutes := setupRoutes(cdrResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest("POST", "/", bytes.NewReader([]byte(`{
			"id": "CDR0003",
			"start_date_time": "2015-06-29T21:39:09Z",
			"stop_date_time": "2015-06-29T21:39:09Z",
			"auth_id": "DE8ACC12E46L89",
			"auth_method": "AUTH_REQUEST",
			"location": {
				"id": "LOC2",
				"type": "UNDERGROUND_GARAGE",
				"name": "Gent Zuid",
				"address": "F.Rooseveltlaan 3A",
				"city": "Gent",
				"postal_code": "9000",
				"country": "BEL",
				"coordinates": {
					"latitude": "50.770774",
					"longitude": "-126.104965"	
				},
				"related_locations": [],
				"evses": [],
				"directions": [{
					"text": "Go Left",
					"language": "en"
				}, {
					"text": "Go Right",
					"language": "en"
				}],
				"facilities": ["BUS_STOP", "TAXI_STAND", "TRAIN_STATION"],
				"charging_when_closed": true,
				"images": [],
				"energy_mix": null,
				"last_updated": "2015-06-29T20:39:09Z"
			},
			"currency": "EUR",
			"tariffs": [],
			"charging_periods": [],
			"total_cost": 4,
			"total_energy": 15.342,
			"total_time": 1.973,
			"last_updated": "2015-06-29T22:01:13Z"
		}`)))

		if err != nil {
			t.Fatal("Creating 'POST /' request failed!")
		}

		cdrRoutes.ServeHTTP(responseRecorder, request)

		cdrParams, err := mockRepository.GetCreateCdrMockData()
		paramsJson, _ := json.Marshal(cdrParams)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"uid": "CDR0003",
			"countryCode": {"String": "", "Valid": false},
			"partyID": {"String": "", "Valid": false},
			"authorizationID": {"String": "", "Valid": false},
			"startDateTime": "2015-06-29T21:39:09Z",
			"stopDateTime": {"Time": "2015-06-29T21:39:09Z", "Valid": true},
			"authID": "DE8ACC12E46L89",
			"authMethod": "AUTH_REQUEST",
			"locationID": 0,
			"meterID": {"String": "", "Valid": false},
			"currency": "EUR",
			"calibrationID": {"Int64": 0, "Valid": false},
			"totalCost": 4,
			"totalEnergy": 15.342,
			"totalTime": 1.973,
			"totalParkingTime": {"Float64": 0, "Valid": false},
			"remark": {"String": "", "Valid": false},
			"lastUpdated": "2015-06-29T22:01:13Z"
		}`))

		locationHeader := responseRecorder.Header().Get("Location")
		expectedLocationHeader := fmt.Sprintf("%s/%s/%s/%s", os.Getenv("API_DOMAIN"), "2.1.1", "cdrs", "CDR0003")

		if locationHeader != expectedLocationHeader {
			t.Fatal("Returned ", locationHeader, " instead of ", expectedLocationHeader)
		}

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"status_code": 1000,
			"status_message": "Success"
		}`), jsondiff.SupersetMatch)
	})
}
