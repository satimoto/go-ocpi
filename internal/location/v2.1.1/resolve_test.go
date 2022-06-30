package location_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/satimoto/go-datastore/pkg/db"
	dbMocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/geom"
	"github.com/satimoto/go-datastore/pkg/util"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/geolocation"
	location "github.com/satimoto/go-ocpi/internal/location/v2.1.1"
	locationMocks "github.com/satimoto/go-ocpi/internal/location/v2.1.1/mocks"
	transportationMocks "github.com/satimoto/go-ocpi/internal/transportation/mocks"
	"github.com/satimoto/go-ocpi/test/mocks"
)

func TestReplaceLocation(t *testing.T) {
	ctx := context.Background()

	t.Run("Create Location", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		locationResolver := locationMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

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
			LastUpdated:        util.ParseTime("2015-03-16T10:10:02Z", nil),
		}

		cred := db.Credential{
			ID:          1,
			CountryCode: "FR",
			PartyID:     "GER",
		}

		countryCode := "DE"
		partyID := "ABC"
		locationResolver.ReplaceLocationByIdentifier(ctx, cred, &countryCode, &partyID, *dto.ID, &dto)

		params, _ := mockRepository.GetCreateLocationMockData()
		paramsJson, _ := json.Marshal(params)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"uid": "LOC1",
			"credentialID": 1,
			"countryCode": {"String": "DE", "Valid": true},
			"partyID": {"String": "ABC", "Valid": true},
			"type": "ON_STREET",
			"name": {"String": "Gent Zuid", "Valid": true},
			"address": "F.Rooseveltlaan 3A",
			"city": "Gent",
			"postalCode": "9000",
			"country": "BEL",
			"geom": {
				"type": "Point",
				"coordinates": [-62.6996, 31.3434]
			},
			"geoLocationID": 1,
			"availableEvses": 0,
			"totalEvses": 0,
			"isRemoteCapable": false,
			"isRfidCapable": false,
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
			"latitudeFloat": 31.3434,
			"longitude": "-62.6996",
			"longitudeFloat": -62.6996
		}`))
	})

	t.Run("Update Location", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		locationResolver := locationMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		locationTypeONSTREET := db.LocationTypeONSTREET
		point, _ := geom.NewPoint("31.3434", "-62.6996")

		mockRepository.SetGetLocationByUidMockData(dbMocks.LocationMockData{
			Location: db.Location{
				Uid:        "LOC1",
				Type:       locationTypeONSTREET,
				Name:       util.SqlNullString("Gent Zuid"),
				Address:    "F.Rooseveltlaan 3A",
				City:       "Gent",
				PostalCode: "9000",
				Country:    "BEL",
				Geom: geom.Geometry4326{
					Coordinates: point,
					Type:        point.GeoJSONType(),
				},
				GeoLocationID:      1,
				ChargingWhenClosed: true,
				LastUpdated:        *util.ParseTime("2015-03-16T10:10:02Z", nil),
			},
		})

		locationTypePARKINGGARAGE := db.LocationTypePARKINGGARAGE

		dto := location.LocationDto{
			Type: &locationTypePARKINGGARAGE,
		}

		cred := db.Credential{
			ID:          1,
			CountryCode: "FR",
			PartyID:     "GER",
		}

		countryCode := "DE"
		partyID := "ABC"
		locationResolver.ReplaceLocationByIdentifier(ctx, cred, &countryCode, &partyID, "LOC1", &dto)

		params, _ := mockRepository.GetUpdateLocationByUidMockData()
		paramsJson, _ := json.Marshal(params)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"uid": "LOC1",
			"countryCode": {"String": "DE", "Valid": true},
			"partyID": {"String": "ABC", "Valid": true},
			"type": "PARKING_GARAGE",
			"name": {"String": "Gent Zuid", "Valid": true},
			"address": "F.Rooseveltlaan 3A",
			"city": "Gent",
			"postalCode": "9000",
			"country": "BEL",
			"geom": {
				"type": "Point",
				"coordinates": [31.3434, -62.6996]
			},
			"geoLocationID": 1,
			"availableEvses": 0,
			"totalEvses": 0,
			"isRemoteCapable": false,
			"isRfidCapable": false,
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
		locationResolver := locationMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		locationTypeONSTREET := db.LocationTypeONSTREET
		point, _ := geom.NewPoint("31.3434", "-62.6996")

		mockRepository.SetGetLocationByUidMockData(dbMocks.LocationMockData{
			Location: db.Location{
				Uid:        "LOC1",
				Type:       locationTypeONSTREET,
				Name:       util.SqlNullString("Gent Zuid"),
				Address:    "F.Rooseveltlaan 3A",
				City:       "Gent",
				PostalCode: "9000",
				Country:    "BEL",
				Geom: geom.Geometry4326{
					Coordinates: point,
					Type:        point.GeoJSONType(),
				},
				GeoLocationID:      1,
				ChargingWhenClosed: true,
				LastUpdated:        *util.ParseTime("2015-03-16T10:10:02Z", nil),
			},
		})

		evseStatusRESERVED := db.EvseStatusRESERVED

		evsesDto := []*evse.EvseDto{}
		evsesDto = append(evsesDto, &evse.EvseDto{
			Uid:               util.NilString("3257"),
			EvseID:            util.NilString("BE-BEC-E04150300-8"),
			Status:            &evseStatusRESERVED,
			PhysicalReference: util.NilString("2"),
			FloorLevel:        util.NilString("-2"),
			LastUpdated:       util.ParseTime("2015-03-16T10:10:02Z", nil),
		})

		dto := location.LocationDto{
			Evses: evsesDto,
		}

		cred := db.Credential{
			ID:          1,
			CountryCode: "FR",
			PartyID:     "GER",
		}

		locationResolver.ReplaceLocation(ctx, cred, "LOC1", &dto)

		params, _ := mockRepository.GetUpdateLocationByUidMockData()
		paramsJson, _ := json.Marshal(params)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"uid": "LOC1",
			"countryCode": {"String": "BE", "Valid": true},
			"partyID": {"String": "BEC", "Valid": true},
			"type": "ON_STREET",
			"name": {"String": "Gent Zuid", "Valid": true},
			"address": "F.Rooseveltlaan 3A",
			"city": "Gent",
			"postalCode": "9000",
			"country": "BEL",
			"geom": {
				"type": "Point",
				"coordinates": [31.3434, -62.6996]
			},
			"geoLocationID": 1,
			"availableEvses": 0,
			"totalEvses": 0,
			"isRemoteCapable": false,
			"isRfidCapable": false,
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
			"evseID": {"String": "BE-BEC-E04150300-8", "Valid": true},
			"locationID": 0,
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
}
