package evse_test

import (
	"context"
	"encoding/json"
	"testing"

	dbMocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
	connector "github.com/satimoto/go-ocpi-api/internal/connector/v2.1.1"
	evse "github.com/satimoto/go-ocpi-api/internal/evse/v2.1.1"
	evseMocks "github.com/satimoto/go-ocpi-api/internal/evse/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/geolocation"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func TestReplaceEvse(t *testing.T) {
	ctx := context.Background()

	t.Run("Create Evse", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		evseResolver := evseMocks.NewResolver(mockRepository)

		evseStatusRESERVED := db.EvseStatusRESERVED

		dto := evse.EvseDto{
			Uid:               util.NilString("3257"),
			EvseID:            util.NilString("BE-BEC-E041503002"),
			Status:            &evseStatusRESERVED,
			PhysicalReference: util.NilString("2"),
			FloorLevel:        util.NilString("-2"),
			LastUpdated:       util.ParseTime("2015-03-16T10:10:02Z", nil),
		}

		evseResolver.ReplaceEvse(ctx, 1, *dto.Uid, &dto)

		params, _ := mockRepository.GetCreateEvseMockData()
		paramsJson, _ := json.Marshal(params)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"uid": "3257",
			"evseID": {"String": "BE-BEC-E041503002", "Valid": true},
			"locationID": 1,
			"status": "RESERVED",
			"geom": {"Geometry4326": {"type": ""}, "Valid": false},
			"geoLocationID": {"Int64": 0, "Valid": false},
			"isRemoteCapable": false,
			"isRfidCapable": false,
			"physicalReference": {"String": "2", "Valid": true},
			"floorLevel": {"String": "-2", "Valid": true},
			"lastUpdated": "2015-03-16T10:10:02Z"
		}`))
	})

	t.Run("Update Evse", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		evseResolver := evseMocks.NewResolver(mockRepository)

		mockRepository.SetGetEvseByUidMockData(dbMocks.EvseMockData{
			Evse: db.Evse{
				Uid:               "3257",
				EvseID:            util.SqlNullString("BE-BEC-E041503002"),
				LocationID:        1,
				Status:            "RESERVED",
				PhysicalReference: util.SqlNullString("2"),
				FloorLevel:        util.SqlNullString("-2"),
				LastUpdated:       *util.ParseTime("2015-03-16T10:10:02Z", nil),
			},
		})

		evseStatusAVAILABLE := db.EvseStatusAVAILABLE

		dto := evse.EvseDto{
			Status: &evseStatusAVAILABLE,
		}

		evseResolver.ReplaceEvse(ctx, 1, "1", &dto)

		params, _ := mockRepository.GetUpdateEvseByUidMockData()
		paramsJson, _ := json.Marshal(params)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"uid": "3257",
			"evseID": {"String": "BE-BEC-E041503002", "Valid": true},
			"status": "AVAILABLE",
			"geom": {"Geometry4326": {"type": ""}, "Valid": false},
			"geoLocationID": {"Int64": 0, "Valid": false},
			"isRemoteCapable": false,
			"isRfidCapable": false,
			"physicalReference": {"String": "2", "Valid": true},
			"floorLevel": {"String": "-2", "Valid": true},
			"lastUpdated": "2015-03-16T10:10:02Z"
		}`))
	})

	t.Run("Update Evse with Connectors", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		evseResolver := evseMocks.NewResolver(mockRepository)

		mockRepository.SetGetEvseByUidMockData(dbMocks.EvseMockData{
			Evse: db.Evse{
				Uid:               "3257",
				EvseID:            util.SqlNullString("BE-BEC-E041503002"),
				LocationID:        1,
				Status:            "RESERVED",
				PhysicalReference: util.SqlNullString("2"),
				FloorLevel:        util.SqlNullString("-2"),
				LastUpdated:       *util.ParseTime("2015-03-16T10:10:02Z", nil),
			},
		})

		mockRepository.SetGetConnectorByUidMockData(dbMocks.ConnectorMockData{
			Connector: db.Connector{
				Uid:         "1",
				EvseID:      1,
				Standard:    "IEC_62196_T2",
				Format:      "CABLE",
				PowerType:   "AC_3_PHASE",
				Voltage:     220,
				Amperage:    16,
				TariffID:    util.SqlNullString("11"),
				LastUpdated: *util.ParseTime("2015-03-16T10:10:02Z", nil),
			},
		})

		connectorsDto := []*connector.ConnectorDto{}
		connectorsDto = append(connectorsDto, &connector.ConnectorDto{
			Id:       util.NilString("1"),
			TariffID: util.NilString("12"),
		})

		dto := evse.EvseDto{
			Connectors: connectorsDto,
		}

		evseResolver.ReplaceEvse(ctx, 1, "1", &dto)

		params, _ := mockRepository.GetUpdateEvseByUidMockData()
		paramsJson, _ := json.Marshal(params)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"uid": "3257",
			"evseID": {"String": "BE-BEC-E041503002", "Valid": true},
			"status": "RESERVED",
			"geom": {"Geometry4326": {"type": ""}, "Valid": false},
			"geoLocationID": {"Int64": 0, "Valid": false},
			"isRemoteCapable": false,
			"isRfidCapable": false,
			"physicalReference": {"String": "2", "Valid": true},
			"floorLevel": {"String": "-2", "Valid": true},
			"lastUpdated": "2015-03-16T10:10:02Z"
		}`))

		connectorParams, _ := mockRepository.GetUpdateConnectorByUidMockData()
		connectorParamsJson, _ := json.Marshal(connectorParams)

		mocks.CompareJson(t, connectorParamsJson, []byte(`{
			"uid": "1",
			"evseID": 1,
			"standard": "IEC_62196_T2",
			"format": "CABLE",
			"powerType": "AC_3_PHASE",
			"voltage": 220,
			"amperage": 16,
			"tariffID": {"String": "12", "Valid": true},
			"termsAndConditions": {"String": "", "Valid": false},
			"lastUpdated": "2015-03-16T10:10:02Z"
		}`))
	})

	t.Run("Update Evse with Coordinates", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		evseResolver := evseMocks.NewResolver(mockRepository)

		mockRepository.SetGetEvseByUidMockData(dbMocks.EvseMockData{
			Evse: db.Evse{
				Uid:               "3257",
				EvseID:            util.SqlNullString("BE-BEC-E041503002"),
				LocationID:        1,
				Status:            "RESERVED",
				PhysicalReference: util.SqlNullString("2"),
				FloorLevel:        util.SqlNullString("-2"),
				LastUpdated:       *util.ParseTime("2015-03-16T10:10:02Z", nil),
			},
		})

		evseStatusAVAILABLE := db.EvseStatusAVAILABLE

		dto := evse.EvseDto{
			Status: &evseStatusAVAILABLE,
			Coordinates: &geolocation.GeoLocationDto{
				Latitude:  "31.3434",
				Longitude: "-62.6996",
			},
		}

		evseResolver.ReplaceEvse(ctx, 1, "1", &dto)

		params, _ := mockRepository.GetUpdateEvseByUidMockData()
		paramsJson, _ := json.Marshal(params)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"uid": "3257",
			"evseID": {"String": "BE-BEC-E041503002", "Valid": true},
			"status": "AVAILABLE",
			"geom": {
				"Geometry4326": {
					"type": "Point",
					"coordinates": [31.3434, -62.6996]
				}, 
				"Valid": true
			},
			"geoLocationID": {"Int64": 1, "Valid": true},
			"isRemoteCapable": false,
			"isRfidCapable": false,
			"physicalReference": {"String": "2", "Valid": true},
			"floorLevel": {"String": "-2", "Valid": true},
			"lastUpdated": "2015-03-16T10:10:02Z"
		}`))

		geolocationParams, _ := mockRepository.GetCreateGeoLocationMockData()
		geolocationParamsJson, _ := json.Marshal(geolocationParams)

		mocks.CompareJson(t, geolocationParamsJson, []byte(`{
			"latitude": "31.3434",
			"latitudeFloat": 31.3434,
			"longitude": "-62.6996",
			"longitudeFloat": -62.6996,
			"name": {"String": "", "Valid": false}
		}`))
	})
}
