package evse_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/satimoto/go-datastore/pkg/db"
	dbMocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	evseMocks "github.com/satimoto/go-ocpi/internal/evse/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi/internal/ocpitype"
	"github.com/satimoto/go-ocpi/test/mocks"
)

func TestReplaceEvse(t *testing.T) {
	ctx := context.Background()

	os.Setenv("RECORD_EVSE_STATUS_PERIODS", "true")
	defer os.Unsetenv("RECORD_EVSE_STATUS_PERIODS")
	
	t.Run("Create Evse", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		evseResolver := evseMocks.NewResolver(mockRepository)

		evseStatusRESERVED := db.EvseStatusRESERVED

		evseDto := dto.EvseDto{
			Uid:               util.NilString("3257"),
			EvseID:            util.NilString("BE-BEC-E041503002"),
			Status:            &evseStatusRESERVED,
			PhysicalReference: util.NilString("2"),
			FloorLevel:        util.NilString("-2"),
			LastUpdated:       ocpitype.ParseOcpiTime("2015-03-16T10:10:02Z", nil),
		}

		evseResolver.ReplaceEvse(ctx, 1, *evseDto.Uid, &evseDto)

		createEvseParams, _ := mockRepository.GetCreateEvseMockData()
		createEvseParamsJson, _ := json.Marshal(createEvseParams)

		mocks.CompareJson(t, createEvseParamsJson, []byte(`{
			"uid": "3257",
			"evseID": {"String": "BE-BEC-E041503002", "Valid": true},
			"identifier": {"String": "BE*BEC*E041503002", "Valid": true},
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

		createEvseStatusPeriodParams, _ := mockRepository.GetCreateEvseStatusPeriodMockData()
		createEvseStatusPeriodParamsJson, _ := json.Marshal(createEvseStatusPeriodParams)

		mocks.CompareJson(t, createEvseStatusPeriodParamsJson, []byte(`{
			"endDate": "0001-01-01T00:00:00Z",
			"evseID": 0,
			"startDate": "0001-01-01T00:00:00Z",
			"status": ""
		}`))
	})

	t.Run("Update Evse", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		evseResolver := evseMocks.NewResolver(mockRepository)

		mockRepository.SetGetEvseByUidMockData(dbMocks.EvseMockData{
			Evse: db.Evse{
				ID:                1,
				Uid:               "3257",
				EvseID:            util.SqlNullString("BE-BEC-E041503002"),
				Identifier:        util.SqlNullString("BE*BEC*E041503002"),
				LocationID:        1,
				Status:            "RESERVED",
				PhysicalReference: util.SqlNullString("2"),
				FloorLevel:        util.SqlNullString("-2"),
				LastUpdated:       *util.ParseTime("2015-03-16T10:00:00Z", nil),
			},
		})

		evseStatusAVAILABLE := db.EvseStatusAVAILABLE

		evseDto := dto.EvseDto{
			Status:      &evseStatusAVAILABLE,
			LastUpdated: ocpitype.ParseOcpiTime("2015-03-16T10:10:02Z", nil),
		}

		evseResolver.ReplaceEvse(ctx, 1, "1", &evseDto)

		params, _ := mockRepository.GetUpdateEvseByUidMockData()
		paramsJson, _ := json.Marshal(params)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"uid": "3257",
			"evseID": {"String": "BE-BEC-E041503002", "Valid": true},
			"identifier": {"String": "BE*BEC*E041503002", "Valid": true},
			"status": "AVAILABLE",
			"geom": {"Geometry4326": {"type": ""}, "Valid": false},
			"geoLocationID": {"Int64": 0, "Valid": false},
			"isRemoteCapable": false,
			"isRfidCapable": false,
			"physicalReference": {"String": "2", "Valid": true},
			"floorLevel": {"String": "-2", "Valid": true},
			"lastUpdated": "2015-03-16T10:10:02Z"
		}`))

		createEvseStatusPeriodParams, _ := mockRepository.GetCreateEvseStatusPeriodMockData()
		createEvseStatusPeriodParamsJson, _ := json.Marshal(createEvseStatusPeriodParams)

		mocks.CompareJson(t, createEvseStatusPeriodParamsJson, []byte(`{
			"endDate": "2015-03-16T10:10:02Z",
			"evseID": 1,
			"startDate": "2015-03-16T10:00:00Z",
			"status": "RESERVED"
		}`))
	})

	t.Run("Update Evse with Connectors", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		evseResolver := evseMocks.NewResolver(mockRepository)

		mockRepository.SetGetEvseByUidMockData(dbMocks.EvseMockData{
			Evse: db.Evse{
				Uid:               "3257",
				EvseID:            util.SqlNullString("BE-BEC-E041503002"),
				Identifier:        util.SqlNullString("BE*BEC*E041503002"),
				LocationID:        1,
				Status:            "RESERVED",
				PhysicalReference: util.SqlNullString("2"),
				FloorLevel:        util.SqlNullString("-2"),
				LastUpdated:       *util.ParseTime("2015-03-16T10:00:00Z", nil),
			},
		})

		mockRepository.SetGetConnectorByEvseMockData(dbMocks.ConnectorMockData{
			Connector: db.Connector{
				Uid:         "1",
				EvseID:      1,
				Identifier:  util.SqlNullString("BE*BEC*E041503002*1"),
				Standard:    "IEC_62196_T2",
				Format:      "CABLE",
				PowerType:   "AC_3_PHASE",
				Voltage:     220,
				Amperage:    16,
				TariffID:    util.SqlNullString("11"),
				LastUpdated: *util.ParseTime("2015-03-16T10:00:00Z", nil),
			},
		})

		connectorsDto := []*dto.ConnectorDto{}
		connectorsDto = append(connectorsDto, &dto.ConnectorDto{
			Id:          util.NilString("1"),
			TariffID:    util.NilString("12"),
			LastUpdated: ocpitype.ParseOcpiTime("2015-03-16T10:10:02Z", nil),
		})

		evseDto := dto.EvseDto{
			Connectors:  connectorsDto,
			LastUpdated: ocpitype.ParseOcpiTime("2015-03-16T10:10:02Z", nil),
		}

		evseResolver.ReplaceEvse(ctx, 1, "1", &evseDto)

		params, _ := mockRepository.GetUpdateEvseByUidMockData()
		paramsJson, _ := json.Marshal(params)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"uid": "3257",
			"evseID": {"String": "BE-BEC-E041503002", "Valid": true},
			"identifier": {"String": "BE*BEC*E041503002", "Valid": true},
			"status": "RESERVED",
			"geom": {"Geometry4326": {"type": ""}, "Valid": false},
			"geoLocationID": {"Int64": 0, "Valid": false},
			"isRemoteCapable": false,
			"isRfidCapable": false,
			"physicalReference": {"String": "2", "Valid": true},
			"floorLevel": {"String": "-2", "Valid": true},
			"lastUpdated": "2015-03-16T10:10:02Z"
		}`))

		connectorParams, _ := mockRepository.GetUpdateConnectorByEvseMockData()
		connectorParamsJson, _ := json.Marshal(connectorParams)

		mocks.CompareJson(t, connectorParamsJson, []byte(`{
			"uid": "1",
			"evseID": 1,
			"identifier": {"String": "BE*BEC*E041503002*1", "Valid": true},
			"standard": "IEC_62196_T2",
			"format": "CABLE",
			"powerType": "AC_3_PHASE",
			"publish": true,
			"voltage": 220,
			"amperage": 16,
			"wattage": 10560,
			"tariffID": {"String": "12", "Valid": true},
			"termsAndConditions": {"String": "", "Valid": false},
			"lastUpdated": "2015-03-16T10:10:02Z"
		}`))

		createEvseStatusPeriodParams, _ := mockRepository.GetCreateEvseStatusPeriodMockData()
		createEvseStatusPeriodParamsJson, _ := json.Marshal(createEvseStatusPeriodParams)

		mocks.CompareJson(t, createEvseStatusPeriodParamsJson, []byte(`{
			"endDate": "0001-01-01T00:00:00Z",
			"evseID": 0,
			"startDate": "0001-01-01T00:00:00Z",
			"status": ""
		}`))
	})

	t.Run("Update Evse with Coordinates", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		evseResolver := evseMocks.NewResolver(mockRepository)

		mockRepository.SetGetEvseByUidMockData(dbMocks.EvseMockData{
			Evse: db.Evse{
				ID:                1,
				Uid:               "3257",
				EvseID:            util.SqlNullString("BE-BEC-E041503002"),
				Identifier:        util.SqlNullString("BE*BEC*E041503002"),
				LocationID:        1,
				Status:            "RESERVED",
				PhysicalReference: util.SqlNullString("2"),
				FloorLevel:        util.SqlNullString("-2"),
				LastUpdated:       *util.ParseTime("2015-03-16T10:00:00Z", nil),
			},
		})

		evseStatusAVAILABLE := db.EvseStatusAVAILABLE

		evseDto := dto.EvseDto{
			Status: &evseStatusAVAILABLE,
			Coordinates: &coreDto.GeoLocationDto{
				Latitude:  "31.3434",
				Longitude: "-62.6996",
			},
			LastUpdated: ocpitype.ParseOcpiTime("2015-03-16T10:10:02Z", nil),
		}

		evseResolver.ReplaceEvse(ctx, 1, "1", &evseDto)

		params, _ := mockRepository.GetUpdateEvseByUidMockData()
		paramsJson, _ := json.Marshal(params)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"uid": "3257",
			"evseID": {"String": "BE-BEC-E041503002", "Valid": true},
			"identifier": {"String": "BE*BEC*E041503002", "Valid": true},
			"status": "AVAILABLE",
			"geom": {
				"Geometry4326": {
					"type": "Point",
					"coordinates": [-62.6996, 31.3434]
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
			"longitudeFloat": -62.6996
		}`))

		createEvseStatusPeriodParams, _ := mockRepository.GetCreateEvseStatusPeriodMockData()
		createEvseStatusPeriodParamsJson, _ := json.Marshal(createEvseStatusPeriodParams)

		mocks.CompareJson(t, createEvseStatusPeriodParamsJson, []byte(`{
			"endDate": "2015-03-16T10:10:02Z",
			"evseID": 1,
			"startDate": "2015-03-16T10:00:00Z",
			"status": "RESERVED"
		}`))
	})
}
