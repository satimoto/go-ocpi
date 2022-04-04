package location_test

import (
	"context"
	"encoding/json"
	"testing"

	dbMocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-datastore/db"
	evse "github.com/satimoto/go-ocpi-api/internal/evse/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/geolocation"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
	locationMocks "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/util"
	"github.com/satimoto/go-ocpi-api/test/mocks"
	"github.com/twpayne/go-geom"
)

func TestReplaceLocation(t *testing.T) {
	ctx := context.Background()

	t.Run("Create Location", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		locationResolver := locationMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

		locationTypeONSTREET := db.LocationTypeONSTREET

		dto := location.LocationDto{
			ID:         util.NilString("LOC1"),
			Type:       &locationTypeONSTREET,
			Name:       util.NilString("Gent Zuid"),
			Address:    util.NilString("F.Rooseveltlaan 3A"),
			City:       util.NilString("Gent"),
			PostalCode: util.NilString("9000"),
			Country:    util.NilString("BEL"),
			Coordinates: &geolocation.GeoLocationDto{
				Latitude:  "31.3434",
				Longitude: "-62.6996",
			},
			ChargingWhenClosed: util.NilBool(true),
			LastUpdated:        util.ParseTime("2015-03-16T10:10:02Z"),
		}

		locationResolver.ReplaceLocation(ctx, *dto.ID, &dto)

		params, _ := mockRepository.GetCreateLocationMockData()
		paramsJson, _ := json.Marshal(params)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"uid": "LOC1",
			"type": "ON_STREET",
			"name": {"String": "Gent Zuid", "Valid": true},
			"address": "F.Rooseveltlaan 3A",
			"city": "Gent",
			"postalCode": "9000",
			"country": "BEL",
			"geom": {},
			"geoLocationID": 1,
			"energyMixID": {"Int64": 0, "Valid": false},
			"openingTimeID": {"Int64": 0, "Valid": false},
			"timeZone": {"String": "", "Valid": false},
			"operatorID": {"Int64": 0, "Valid": false},
			"ownerID": {"Int64": 0, "Valid": false},
			"suboperatorID": {"Int64": 0, "Valid": false},
			"chargingWhenClosed": true,
			"lastUpdated": "2015-03-16T10:10:02Z"
		}`))

		geolocationParams, _ := mockRepository.GetCreateGeoLocationMockData()
		geolocationParamsJson, _ := json.Marshal(geolocationParams)

		mocks.CompareJson(t, geolocationParamsJson, []byte(`{
			"latitude": "31.3434",
			"longitude": "-62.6996",
			"name": {"String": "", "Valid": false}
		}`))
	})

	t.Run("Update Location", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		locationResolver := locationMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

		locationTypeONSTREET := db.LocationTypeONSTREET
		point, _ := util.NewGeomPoint("31.3434", "-62.6996")

		mockRepository.SetGetLocationByUidMockData(dbMocks.LocationMockData{
			Location: db.Location{
				Uid:                "LOC1",
				Type:               locationTypeONSTREET,
				Name:               util.SqlNullString("Gent Zuid"),
				Address:            "F.Rooseveltlaan 3A",
				City:               "Gent",
				PostalCode:         "9000",
				Country:            "BEL",
				Geom:               *geom.NewPointFlat(geom.XY, point),
				GeoLocationID:      1,
				ChargingWhenClosed: true,
				LastUpdated:        *util.ParseTime("2015-03-16T10:10:02Z"),
			},
		})

		locationTypePARKINGGARAGE := db.LocationTypePARKINGGARAGE

		dto := location.LocationDto{
			Type: &locationTypePARKINGGARAGE,
		}

		locationResolver.ReplaceLocation(ctx, "LOC1", &dto)

		params, _ := mockRepository.GetUpdateLocationByUidMockData()
		paramsJson, _ := json.Marshal(params)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"uid": "LOC1",
			"type": "PARKING_GARAGE",
			"name": {"String": "Gent Zuid", "Valid": true},
			"address": "F.Rooseveltlaan 3A",
			"city": "Gent",
			"postalCode": "9000",
			"country": "BEL",
			"geom": {},
			"geoLocationID": 1,
			"energyMixID": {"Int64": 0, "Valid": false},
			"openingTimeID": {"Int64": 0, "Valid": false},
			"timeZone": {"String": "", "Valid": false},
			"operatorID": {"Int64": 0, "Valid": false},
			"ownerID": {"Int64": 0, "Valid": false},
			"suboperatorID": {"Int64": 0, "Valid": false},
			"chargingWhenClosed": true,
			"lastUpdated": "2015-03-16T10:10:02Z"
		}`))
	})

	t.Run("Update Location with Evses", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		locationResolver := locationMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

		locationTypeONSTREET := db.LocationTypeONSTREET
		point, _ := util.NewGeomPoint("31.3434", "-62.6996")

		mockRepository.SetGetLocationByUidMockData(dbMocks.LocationMockData{
			Location: db.Location{
				Uid:                "LOC1",
				Type:               locationTypeONSTREET,
				Name:               util.SqlNullString("Gent Zuid"),
				Address:            "F.Rooseveltlaan 3A",
				City:               "Gent",
				PostalCode:         "9000",
				Country:            "BEL",
				Geom:               *geom.NewPointFlat(geom.XY, point),
				GeoLocationID:      1,
				ChargingWhenClosed: true,
				LastUpdated:        *util.ParseTime("2015-03-16T10:10:02Z"),
			},
		})

		evseStatusRESERVED := db.EvseStatusRESERVED

		evsesDto := []*evse.EvseDto{}
		evsesDto = append(evsesDto, &evse.EvseDto{
			Uid:               util.NilString("3257"),
			EvseID:            util.NilString("BE-BEC-E041503002"),
			Status:            &evseStatusRESERVED,
			PhysicalReference: util.NilString("2"),
			FloorLevel:        util.NilString("-2"),
			LastUpdated:       util.ParseTime("2015-03-16T10:10:02Z"),
		})

		dto := location.LocationDto{
			Evses: evsesDto,
		}

		locationResolver.ReplaceLocation(ctx, "LOC1", &dto)

		params, _ := mockRepository.GetUpdateLocationByUidMockData()
		paramsJson, _ := json.Marshal(params)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"uid": "LOC1",
			"type": "ON_STREET",
			"name": {"String": "Gent Zuid", "Valid": true},
			"address": "F.Rooseveltlaan 3A",
			"city": "Gent",
			"postalCode": "9000",
			"country": "BEL",
			"geom": {},
			"geoLocationID": 1,
			"energyMixID": {"Int64": 0, "Valid": false},
			"openingTimeID": {"Int64": 0, "Valid": false},
			"timeZone": {"String": "", "Valid": false},
			"operatorID": {"Int64": 0, "Valid": false},
			"ownerID": {"Int64": 0, "Valid": false},
			"suboperatorID": {"Int64": 0, "Valid": false},
			"chargingWhenClosed": true,
			"lastUpdated": "2015-03-16T10:10:02Z"
		}`))

		evseParams, _ := mockRepository.GetCreateEvseMockData()
		evseParamsJson, _ := json.Marshal(evseParams)

		mocks.CompareJson(t, evseParamsJson, []byte(`{
			"uid": "3257",
			"evseID": {"String": "BE-BEC-E041503002", "Valid": true},
			"locationID": 0,
			"status": "RESERVED",
			"geom": null,
			"geoLocationID": {"Int64": 0, "Valid": false},
			"physicalReference": {"String": "2", "Valid": true},
			"floorLevel": {"String": "-2", "Valid": true},
			"lastUpdated": "2015-03-16T10:10:02Z"
		}`))
	})
}
