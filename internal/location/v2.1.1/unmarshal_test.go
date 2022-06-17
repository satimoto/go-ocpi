package location_test

import (
	"encoding/json"
	"testing"

	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func TestLocationUnmarshal(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		dto := location.LocationDto{}
		response := []byte(`{
			"id": null,
			"type": null,
			"address": null,
			"city": null,
			"postal_code": null,
			"country": null,
			"coordinates": null,
			"related_locations": null,
			"evses": null,
			"directions": null,
			"facilities": null,
			"charging_when_closed": null,
			"images": null,
			"energy_mix": null,
			"last_updated": null
		}`)

		json.Unmarshal([]byte(`{}`), &dto)
		responseJson, _ := json.Marshal(dto)

		mocks.CompareJson(t, responseJson, response)
	})

	t.Run("With Evse", func(t *testing.T) {
		dto := location.LocationDto{}
		request := []byte(`{
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
		}`)

		json.Unmarshal(request, &dto)
		responseJson, _ := json.Marshal(dto)

		mocks.CompareJson(t, responseJson, request)
	})

	t.Run("With Directions, Facilities", func(t *testing.T) {
		dto := location.LocationDto{}
		request := []byte(`{
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
		}`)

		json.Unmarshal(request, &dto)
		responseJson, _ := json.Marshal(dto)

		mocks.CompareJson(t, responseJson, request)
	})

	t.Run("With Coordinates, Related locations, Images", func(t *testing.T) {
		dto := location.LocationDto{}
		request := []byte(`{
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
		}`)

		json.Unmarshal(request, &dto)
		responseJson, _ := json.Marshal(dto)

		mocks.CompareJson(t, responseJson, request)
	})
}
