package location_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"

	dbMocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-datastore/db"
	locationMocks "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/util"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func TestCreateLocationDto(t *testing.T) {
	ctx := context.Background()

	t.Run("Empty", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		locationResolver := locationMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

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
			"related_locations": [],
			"evses": [],
			"directions": [],
			"facilities": [],
			"charging_when_closed": false,
			"images": [],
			"energy_mix": null,
			"last_updated": "0001-01-01T00:00:00Z"
		}`))
	})

	t.Run("With Evse", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		locationResolver := locationMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

		statusSchedules := []db.StatusSchedule{}
		statusSchedules = append(statusSchedules, db.StatusSchedule{
			PeriodBegin: *util.ParseTime("2018-12-16T10:10:02Z"),
			PeriodEnd:   util.SqlNullTime(*util.ParseTime("2018-12-16T10:30:02Z")),
			Status:      db.EvseStatusBLOCKED,
		})
		statusSchedules = append(statusSchedules, db.StatusSchedule{
			PeriodBegin: *util.ParseTime("2018-12-16T10:30:02Z"),
			PeriodEnd:   util.SqlNullTime(*util.ParseTime("2018-12-16T11:00:02Z")),
			Status:      db.EvseStatusCHARGING,
		})
		statusSchedules = append(statusSchedules, db.StatusSchedule{
			PeriodBegin: *util.ParseTime("2018-12-16T11:00:02Z"),
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
			TariffID:    sql.NullString{String: "11"},
			LastUpdated: *util.ParseTime("2015-03-16T10:10:02Z"),
		})
		connectors = append(connectors, db.Connector{
			Uid:         "2",
			Standard:    "IEC_62196_T2_COMBO",
			Format:      "SOCKET",
			PowerType:   "AC_1_PHASE",
			Voltage:     110,
			Amperage:    32,
			TariffID:    sql.NullString{String: "9"},
			LastUpdated: *util.ParseTime("2015-03-18T08:12:01Z"),
		})
		mockRepository.SetListConnectorsMockData(dbMocks.ConnectorsMockData{Connectors: connectors, Error: nil})

		evses := []db.Evse{}
		evses = append(evses, db.Evse{
			Uid:               "3257",
			EvseID:            sql.NullString{String: "BE-BEC-E041503002"},
			Status:            db.EvseStatusRESERVED,
			PhysicalReference: sql.NullString{String: "2"},
			FloorLevel:        sql.NullString{String: "-2"},
			LastUpdated:       *util.ParseTime("2015-06-29T20:39:09Z"),
		})
		mockRepository.SetListEvsesMockData(dbMocks.EvsesMockData{Evses: evses, Error: nil})

		operator := db.BusinessDetail{
			Name: "BeCharged",
		}
		mockRepository.SetGetBusinessDetailMockData(dbMocks.BusinessDetailMockData{BusinessDetail: operator, Error: nil})

		loc := db.Location{
			Uid:         "LOC1",
			Type:        db.LocationTypeONSTREET,
			Name:        sql.NullString{String: "Gent Zuid"},
			Address:     "F.Rooseveltlaan 3A",
			City:        "Gent",
			PostalCode:  "9000",
			Country:     "BEL",
			OperatorID:  sql.NullInt64{Int64: 10, Valid: true},
			LastUpdated: *util.ParseTime("2015-06-29T20:39:09Z"),
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
			"related_locations": [],
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
				"directions": [],
				"parking_restrictions": [],
				"images": [],
				"last_updated": "2015-06-29T20:39:09Z"
			}],
			"directions": [],
			"operator": {
				"name": "BeCharged"
			},
			"facilities": [],
			"charging_when_closed": false,
			"images": [],
			"energy_mix": null,
			"last_updated": "2015-06-29T20:39:09Z"
		}`))
	})

	t.Run("With Directions, Facilities", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		locationResolver := locationMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

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
			Name:               sql.NullString{String: "Gent Zuid"},
			Address:            "F.Rooseveltlaan 3A",
			City:               "Gent",
			PostalCode:         "9000",
			Country:            "BEL",
			ChargingWhenClosed: true,
			LastUpdated:        *util.ParseTime("2015-06-29T20:39:09Z"),
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
		}`))
	})

	t.Run("With Coordinates, Related locations, Images", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		locationResolver := locationMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

		coordinates := db.GeoLocation{
			Latitude:  "50.770774",
			Longitude: "-126.104965",
		}
		mockRepository.SetGetGeoLocationMockData(dbMocks.GeoLocationMockData{GeoLocation: coordinates, Error: nil})

		relatedLocations := []db.GeoLocation{}
		relatedLocations = append(relatedLocations, db.GeoLocation{
			Latitude:  "50.770773",
			Longitude: "-126.104966",
			Name:      sql.NullString{String: "Entrance"},
		})
		relatedLocations = append(relatedLocations, db.GeoLocation{
			Latitude:  "50.77077443",
			Longitude: "-126.104963",
			Name:      sql.NullString{String: "Exit"},
		})
		mockRepository.SetListRelatedLocationsMockData(dbMocks.GeoLocationsMockData{GeoLocations: relatedLocations, Error: nil})

		images := []db.Image{}
		images = append(images, db.Image{
			Url:      "https://business.com/logo.png",
			Category: db.ImageCategoryOPERATOR,
			Type:     "png",
		})
		images = append(images, db.Image{
			Url:       "https://business.com/logo2.jpg",
			Thumbnail: sql.NullString{String: "https://business.com/logo2-thumb.jpg"},
			Category:  db.ImageCategoryENTRANCE,
			Type:      "jpg",
			Width:     sql.NullInt32{Int32: 180},
			Height:    sql.NullInt32{Int32: 180},
		})
		mockRepository.SetListLocationImagesMockData(dbMocks.ImagesMockData{Images: images, Error: nil})

		loc := db.Location{
			Uid:                "LOC2",
			Type:               db.LocationTypePARKINGLOT,
			Name:               sql.NullString{String: "Gent Zuid"},
			Address:            "F.Rooseveltlaan 3A",
			City:               "Gent",
			PostalCode:         "9000",
			Country:            "BEL",
			GeoLocationID:      10,
			ChargingWhenClosed: true,
			LastUpdated:        *util.ParseTime("2015-06-29T20:39:09Z"),
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
				"longitude": "-126.104966",
				"name": "Entrance"
			}, {
				"latitude": "50.77077443",
				"longitude": "-126.104963",
				"name": "Exit"
			}],
			"evses": [],
			"directions": [],
			"facilities": [],
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
			"energy_mix": null,
			"last_updated": "2015-06-29T20:39:09Z"
		}`))
	})
}
