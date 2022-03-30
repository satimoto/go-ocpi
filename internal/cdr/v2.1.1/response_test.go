package cdr_test

import (
	"context"
	"encoding/json"
	"testing"

	dbMocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-datastore/db"
	cdrMocks "github.com/satimoto/go-ocpi-api/internal/cdr/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/util"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func TestCreateCdrPayload(t *testing.T) {
	ctx := context.Background()

	t.Run("Empty", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		cdrResolver := cdrMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

		sess := db.Cdr{
			ID: 1,
		}

		response := cdrResolver.CreateCdrPayload(ctx, sess)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"id": "",
			"start_date_time": "0001-01-01T00:00:00Z",
			"auth_id": "",
			"auth_method": "",
			"location": null,
			"currency": "",
			"tariffs": [],
			"charging_periods": [],
			"total_cost": 0,
			"total_energy": 0,
			"total_time": 0,
			"last_updated": "0001-01-01T00:00:00Z"
		}`))
	})

	t.Run("Simple", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		cdrResolver := cdrMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

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

		response := cdrResolver.CreateCdrPayload(ctx, c)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
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
		}`))
	})

	t.Run("With charge periods", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		cdrResolver := cdrMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

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

		sess := db.Cdr{
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

		response := cdrResolver.CreateCdrPayload(ctx, sess)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
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
	}`))
	})
}
