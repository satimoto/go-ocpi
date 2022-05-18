package session_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/satimoto/go-datastore/pkg/db"
	dbMocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/util"
	sessionMocks "github.com/satimoto/go-ocpi-api/internal/session/v2.1.1/mocks"
	transportationMocks "github.com/satimoto/go-ocpi-api/internal/transportation/mocks"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func TestCreateSessionDto(t *testing.T) {
	ctx := context.Background()

	t.Run("Empty", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		sessionResolver := sessionMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		sess := db.Session{}

		response := sessionResolver.CreateSessionDto(ctx, sess)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"id": "",
			"start_datetime": "0001-01-01T00:00:00Z",
			"kwh": 0,
			"auth_id": "",
			"auth_method": "",
			"location": null,
			"currency": "",
			"charging_periods": [],
			"status": "",
			"last_updated": "0001-01-01T00:00:00Z"
		}`))
	})

	t.Run("Simple", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		sessionResolver := sessionMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

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

		response := sessionResolver.CreateSessionDto(ctx, sess)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
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
		}`))
	})

	t.Run("With charge periods", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		sessionResolver := sessionMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

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

		response := sessionResolver.CreateSessionDto(ctx, sess)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
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
		}`))
	})
}
