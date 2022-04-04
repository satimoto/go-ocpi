package evse_test

import (
	"encoding/json"
	"testing"

	evse "github.com/satimoto/go-ocpi-api/internal/evse/v2.1.1"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func TestEvseUnmarshal(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		dto := evse.EvseDto{}
		request := []byte(`{}`)

		json.Unmarshal(request, &dto)
		responseJson, _ := json.Marshal(dto)

		mocks.CompareJson(t, responseJson, request)
	})

	t.Run("With Status Schedules", func(t *testing.T) {
		dto := evse.EvseDto{}
		request := []byte(`{
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
			"capabilities": [],
			"connectors": [],
			"physical_reference": "2",
			"floor_level": "-2",
			"directions": [],
			"parking_restrictions": [],
			"images": [],
			"last_updated": "2015-06-29T20:39:09Z"
		}`)

		json.Unmarshal(request, &dto)
		responseJson, _ := json.Marshal(dto)

		mocks.CompareJson(t, responseJson, request)
	})

	t.Run("With Capabilities", func(t *testing.T) {
		dto := evse.EvseDto{}
		request := []byte(`{
			"uid": "3257",
			"evse_id": "BE-BEC-E041503002",
			"status": "RESERVED",
			"status_schedule": [],
			"capabilities": ["RESERVABLE"],
			"connectors": [],
			"directions": [],
			"parking_restrictions": [],
			"images": [],
			"last_updated": "2015-06-29T20:39:09Z"
		}`)

		json.Unmarshal(request, &dto)
		responseJson, _ := json.Marshal(dto)

		mocks.CompareJson(t, responseJson, request)
	})

	t.Run("With Connectors", func(t *testing.T) {
		dto := evse.EvseDto{}
		request := []byte(`{
			"uid": "3256",
			"evse_id": "BE-BEC-E041503001",
			"status": "AVAILABLE",
			"status_schedule": [],
			"capabilities": [],
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
			"physical_reference": "1",
			"floor_level": "-1",
			"directions": [],
			"parking_restrictions": [],
			"images": [],
			"last_updated": "2015-06-28T08:12:01Z"
		}`)

		json.Unmarshal(request, &dto)
		responseJson, _ := json.Marshal(dto)

		mocks.CompareJson(t, responseJson, request)
	})
}
