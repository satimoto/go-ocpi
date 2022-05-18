package cdr_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi"
	"github.com/nsf/jsondiff"
	"github.com/satimoto/go-datastore/pkg/db"
	dbMocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/util"
	cdr "github.com/satimoto/go-ocpi-api/internal/cdr/v2.1.1"
	cdrMocks "github.com/satimoto/go-ocpi-api/internal/cdr/v2.1.1/mocks"
	transportationMocks "github.com/satimoto/go-ocpi-api/internal/transportation/mocks"
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
		cdrResolver := cdrMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))
		cdrRoutes := setupRoutes(cdrResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodGet, "/", nil)

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
		cdrResolver := cdrMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))
		cdrRoutes := setupRoutes(cdrResolver)
		responseRecorder := httptest.NewRecorder()

		c := db.Cdr{
			ID:            1,
			Uid:           "CDR0001",
			StartDateTime: *util.ParseTime("2015-06-29T21:39:09Z", nil),
			StopDateTime:  util.SqlNullTime(util.ParseTime("2015-06-29T21:39:09Z", nil)),
			AuthID:        "DE8ACC12E46L89",
			AuthMethod:    db.AuthMethodTypeAUTHREQUEST,
			Currency:      "EUR",
			TotalCost:     4.0,
			TotalEnergy:   15.342,
			TotalTime:     1.973,
			LastUpdated:   *util.ParseTime("2015-06-29T22:01:13Z", nil),
		}
		mockRepository.SetGetCdrByUidMockData(dbMocks.CdrMockData{Cdr: c, Error: nil})

		request, err := http.NewRequest(http.MethodGet, "/CDR0001", nil)

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
		cdrResolver := cdrMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))
		cdrRoutes := setupRoutes(cdrResolver)
		responseRecorder := httptest.NewRecorder()

		chargingPeriods := []db.ChargingPeriod{}
		chargingPeriods = append(chargingPeriods, db.ChargingPeriod{
			StartDateTime: *util.ParseTime("2015-06-29T22:39:09Z", nil),
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
			StartDateTime:    *util.ParseTime("2015-06-29T21:39:09Z", nil),
			StopDateTime:     util.SqlNullTime(util.ParseTime("2015-06-29T21:39:09Z", nil)),
			AuthID:           "DE8ACC12E46L89",
			AuthMethod:       db.AuthMethodTypeAUTHREQUEST,
			Currency:         "EUR",
			TotalCost:        4.0,
			TotalEnergy:      15.342,
			TotalTime:        1.973,
			TotalParkingTime: util.SqlNullFloat64(45),
			LastUpdated:      *util.ParseTime("2015-06-29T22:01:13Z", nil),
		}
		mockRepository.SetGetCdrByUidMockData(dbMocks.CdrMockData{Cdr: c, Error: nil})

		request, err := http.NewRequest(http.MethodGet, "/CDR0002", nil)

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
		cdrResolver := cdrMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))
		cdrRoutes := setupRoutes(cdrResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(`{
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
				"evses": [{
                    "uid": "9785",
                    "status": "UNKNOWN",
                    "capabilities": [
                        "CHARGING_PROFILE_CAPABLE",
                        "REMOTE_START_STOP_CAPABLE",
                        "RFID_READER"
                    ],
                    "connectors": [{
                        "id": "9785",
                        "standard": "IEC_62196_T2",
                        "format": "SOCKET",
                        "power_type": "AC_3_PHASE",
                        "voltage": 230,
                        "amperage": 16,
                        "last_updated": "2022-04-07T07:19:49Z",
                        "max_power": 11040
                    }],
                    "last_updated": "2022-04-07T03:09:30Z",
                    "physical_reference": "AL106",
                    "evse_id": "NL*EVN*E30332*9785",
                    "parking_restrictions": [
                        "EV_ONLY"
                    ]
                }],
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

		requestCtx := request.Context()
		requestCtx = context.WithValue(requestCtx, "credential", &db.Credential{ID: 1, CountryCode: "FR", PartyID: "GER"})
		cdrRoutes.ServeHTTP(responseRecorder, request.WithContext(requestCtx))

		cdrParams, err := mockRepository.GetCreateCdrMockData()
		paramsJson, _ := json.Marshal(cdrParams)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"uid": "CDR0003",
			"credentialID": 1,
			"countryCode": {"String": "NL", "Valid": true},
			"partyID": {"String": "EVN", "Valid": true},
			"authorizationID": {"String": "", "Valid": false},
			"startDateTime": "2015-06-29T21:39:09Z",
			"stopDateTime": {"Time": "2015-06-29T21:39:09Z", "Valid": true},
			"authID": "DE8ACC12E46L89",
			"authMethod": "AUTH_REQUEST",
			"userID": 0,
			"tokenID": 0,
			"locationID": 0,
			"evseID": 0,
			"connectorID": 0,
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
