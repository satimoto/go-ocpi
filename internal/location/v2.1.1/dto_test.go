package location_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/satimoto/go-datastore/pkg/db"
	dbMocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/util"
	locationMocks "github.com/satimoto/go-ocpi/internal/location/v2.1.1/mocks"
	transportationMocks "github.com/satimoto/go-ocpi/internal/transportation/mocks"
	"github.com/satimoto/go-ocpi/test/mocks"
)

func TestCreateLocationDto(t *testing.T) {
	ctx := context.Background()

	t.Run("Empty", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		locationResolver := locationMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		loc := db.Location{}

		response := locationResolver.CreateLocationDto(ctx, loc)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"id": "",
			"type": "",
			"address": "",
			"city": "",
			"postal_code": "",
			"country": "",
			"coordinates": null,
			"charging_when_closed": false,
			"last_updated": null
		}`))
	})

	t.Run("With Evse", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		locationResolver := locationMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		statusSchedules := []db.StatusSchedule{}
		statusSchedules = append(statusSchedules, db.StatusSchedule{
			PeriodBegin: *util.ParseTime("2018-12-16T10:10:02Z", nil),
			PeriodEnd:   util.SqlNullTime(*util.ParseTime("2018-12-16T10:30:02Z", nil)),
			Status:      db.EvseStatusBLOCKED,
		})
		statusSchedules = append(statusSchedules, db.StatusSchedule{
			PeriodBegin: *util.ParseTime("2018-12-16T10:30:02Z", nil),
			PeriodEnd:   util.SqlNullTime(*util.ParseTime("2018-12-16T11:00:02Z", nil)),
			Status:      db.EvseStatusCHARGING,
		})
		statusSchedules = append(statusSchedules, db.StatusSchedule{
			PeriodBegin: *util.ParseTime("2018-12-16T11:00:02Z", nil),
			Status:      db.EvseStatusAVAILABLE,
		})
		mockRepository.SetListStatusSchedulesMockData(dbMocks.StatusSchedulesMockData{StatusSchedules: statusSchedules, Error: nil})

		capabilities := []db.Capability{}
		capabilities = append(capabilities, db.Capability{
			Text: "RESERVABLE",
		})
		mockRepository.SetListEvseCapabilitiesMockData(dbMocks.CapabilitiesMockData{Capabilities: capabilities, Error: nil})

		connectors := []db.Connector{}
		connectors = append(connectors, db.Connector{
			Uid:         "1",
			Standard:    "IEC_62196_T2",
			Format:      "CABLE",
			PowerType:   "AC_3_PHASE",
			Voltage:     220,
			Amperage:    16,
			TariffID:    util.SqlNullString("11"),
			LastUpdated: *util.ParseTime("2015-03-16T10:10:02Z", nil),
		})
		connectors = append(connectors, db.Connector{
			Uid:         "2",
			Standard:    "IEC_62196_T2_COMBO",
			Format:      "SOCKET",
			PowerType:   "AC_1_PHASE",
			Voltage:     110,
			Amperage:    32,
			TariffID:    util.SqlNullString("9"),
			LastUpdated: *util.ParseTime("2015-03-18T08:12:01Z", nil),
		})
		mockRepository.SetListConnectorsMockData(dbMocks.ConnectorsMockData{Connectors: connectors, Error: nil})

		evses := []db.Evse{}
		evses = append(evses, db.Evse{
			Uid:               "3257",
			EvseID:            util.SqlNullString("BE-BEC-E041503002"),
			Status:            db.EvseStatusRESERVED,
			PhysicalReference: util.SqlNullString("2"),
			FloorLevel:        util.SqlNullString("-2"),
			LastUpdated:       *util.ParseTime("2015-06-29T20:39:09Z", nil),
		})
		mockRepository.SetListEvsesMockData(dbMocks.EvsesMockData{Evses: evses, Error: nil})

		operator := db.BusinessDetail{
			Name: "BeCharged",
		}
		mockRepository.SetGetBusinessDetailMockData(dbMocks.BusinessDetailMockData{BusinessDetail: operator, Error: nil})

		loc := db.Location{
			Uid:         "LOC1",
			Type:        db.LocationTypeONSTREET,
			Name:        util.SqlNullString("Gent Zuid"),
			Address:     "F.Rooseveltlaan 3A",
			City:        "Gent",
			PostalCode:  "9000",
			Country:     "BEL",
			OperatorID:  sql.NullInt64{Int64: 10, Valid: true},
			LastUpdated: *util.ParseTime("2015-06-29T20:39:09Z", nil),
		}

		response := locationResolver.CreateLocationDto(ctx, loc)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"id": "LOC1",
			"type": "ON_STREET",
			"name": "Gent Zuid",
			"address": "F.Rooseveltlaan 3A",
			"city": "Gent",
			"postal_code": "9000",
			"country": "BEL",
			"coordinates": null,
			"evses": [{
				"uid": "3257",
				"evse_id": "BE-BEC-E041503002",
				"status": "RESERVED",
				"status_schedule": [{
					"period_begin": "2018-12-16T10:10:02Z",
					"period_end": "2018-12-16T10:30:02Z",
					"status": "BLOCKED"
				}, {
					"period_begin": "2018-12-16T10:30:02Z",
					"period_end": "2018-12-16T11:00:02Z",
					"status": "CHARGING"
				}, {
					"period_begin": "2018-12-16T11:00:02Z",
					"status": "AVAILABLE"
				}],
				"capabilities": ["RESERVABLE"],
				"connectors": [{
					"id": "1",
					"standard": "IEC_62196_T2",
					"format": "CABLE",
					"power_type": "AC_3_PHASE",
					"voltage": 220,
					"amperage": 16,
					"tariff_id": "11",
					"last_updated": "2015-03-16T10:10:02Z"
				}, {
					"id": "2",
					"standard": "IEC_62196_T2_COMBO",
					"format": "SOCKET",
					"power_type": "AC_1_PHASE",
					"voltage": 110,
					"amperage": 32,
					"tariff_id": "9",
					"last_updated": "2015-03-18T08:12:01Z"
				}],
				"physical_reference": "2",
				"floor_level": "-2",
				"last_updated": "2015-06-29T20:39:09Z"
			}],
			"operator": {
				"name": "BeCharged"
			},
			"charging_when_closed": false,
			"last_updated": "2015-06-29T20:39:09Z"
		}`))
	})

	t.Run("With Directions, Facilities", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		locationResolver := locationMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		directions := []db.DisplayText{}
		directions = append(directions, db.DisplayText{
			Text:     "Go Left",
			Language: "en",
		})
		directions = append(directions, db.DisplayText{
			Text:     "Go Right",
			Language: "en",
		})
		mockRepository.SetListLocationDirectionsMockData(dbMocks.DisplayTextsMockData{DisplayTexts: directions, Error: nil})

		facilities := []db.Facility{}
		facilities = append(facilities, db.Facility{
			Text: "BUS_STOP",
		})
		facilities = append(facilities, db.Facility{
			Text: "TAXI_STAND",
		})
		facilities = append(facilities, db.Facility{
			Text: "TRAIN_STATION",
		})
		mockRepository.SetListLocationFacilitiesMockData(dbMocks.FacilitiesMockData{Facilities: facilities, Error: nil})

		loc := db.Location{
			Uid:                "LOC2",
			Type:               db.LocationTypeUNDERGROUNDGARAGE,
			Name:               util.SqlNullString("Gent Zuid"),
			Address:            "F.Rooseveltlaan 3A",
			City:               "Gent",
			PostalCode:         "9000",
			Country:            "BEL",
			ChargingWhenClosed: true,
			LastUpdated:        *util.ParseTime("2015-06-29T20:39:09Z", nil),
		}

		response := locationResolver.CreateLocationDto(ctx, loc)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"id": "LOC2",
			"type": "UNDERGROUND_GARAGE",
			"name": "Gent Zuid",
			"address": "F.Rooseveltlaan 3A",
			"city": "Gent",
			"postal_code": "9000",
			"country": "BEL",
			"coordinates": null,
			"directions": [{
				"text": "Go Left",
				"language": "en"
			}, {
				"text": "Go Right",
				"language": "en"
			}],
			"facilities": ["BUS_STOP", "TAXI_STAND", "TRAIN_STATION"],
			"charging_when_closed": true,
			"last_updated": "2015-06-29T20:39:09Z"
		}`))
	})

	t.Run("With Coordinates, Related locations, Images", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		locationResolver := locationMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		coordinates := db.GeoLocation{
			Latitude:  "50.770774",
			Longitude: "-126.104965",
		}
		mockRepository.SetGetGeoLocationMockData(dbMocks.GeoLocationMockData{GeoLocation: coordinates, Error: nil})

		relatedLocations := []db.AdditionalGeoLocation{}
		relatedLocations = append(relatedLocations, db.AdditionalGeoLocation{
			Latitude:  "50.770773",
			Longitude: "-126.104966",
		})
		relatedLocations = append(relatedLocations, db.AdditionalGeoLocation{
			Latitude:      "50.77077443",
			Longitude:     "-126.104963",
			DisplayTextID: util.SqlNullInt64(2),
		})
		mockRepository.SetGetDisplayTextMockData(dbMocks.DisplayTextMockData{DisplayText: db.DisplayText{
			Language: "nl",
			Text:     "Bloemenspeciaalzaak Bergmans (Store)",
		}})
		mockRepository.SetListAdditionalGeoLocationsMockData(dbMocks.AdditionalGeoLocationsMockData{AdditionalGeoLocations: relatedLocations, Error: nil})

		images := []db.Image{}
		images = append(images, db.Image{
			Url:      "https://business.com/logo.png",
			Category: db.ImageCategoryOPERATOR,
			Type:     "png",
		})
		images = append(images, db.Image{
			Url:       "https://business.com/logo2.jpg",
			Thumbnail: util.SqlNullString("https://business.com/logo2-thumb.jpg"),
			Category:  db.ImageCategoryENTRANCE,
			Type:      "jpg",
			Width:     sql.NullInt32{Int32: 180},
			Height:    sql.NullInt32{Int32: 180},
		})
		mockRepository.SetListLocationImagesMockData(dbMocks.ImagesMockData{Images: images, Error: nil})

		loc := db.Location{
			Uid:                "LOC2",
			Type:               db.LocationTypePARKINGLOT,
			Name:               util.SqlNullString("Gent Zuid"),
			Address:            "F.Rooseveltlaan 3A",
			City:               "Gent",
			PostalCode:         "9000",
			Country:            "BEL",
			GeoLocationID:      10,
			ChargingWhenClosed: true,
			LastUpdated:        *util.ParseTime("2015-06-29T20:39:09Z", nil),
		}

		response := locationResolver.CreateLocationDto(ctx, loc)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"id": "LOC2",
			"type": "PARKING_LOT",
			"name": "Gent Zuid",
			"address": "F.Rooseveltlaan 3A",
			"city": "Gent",
			"postal_code": "9000",
			"country": "BEL",
			"coordinates": {
				"latitude": "50.770774",
				"longitude": "-126.104965"
			},
			"related_locations": [{
				"latitude": "50.770773",
				"longitude": "-126.104966"
			}, {
				"latitude": "50.77077443",
				"longitude": "-126.104963",
				"name": {
					"language": "nl",
					"text": "Bloemenspeciaalzaak Bergmans (Store)"
				}
			}],
			"charging_when_closed": true,
			"images": [{
				"url": "https://business.com/logo.png",
				"category": "OPERATOR",
				"type": "png"
			}, {
				"url": "https://business.com/logo2.jpg",
				"thumbnail": "https://business.com/logo2-thumb.jpg",
				"category": "ENTRANCE",
				"type": "jpg",
				"width": 180,
				"height": 180
			}],
			"last_updated": "2015-06-29T20:39:09Z"
		}`))
	})
}
