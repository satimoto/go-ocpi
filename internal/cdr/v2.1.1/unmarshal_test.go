package cdr_test

import (
	"encoding/json"
	"testing"

	cdr "github.com/satimoto/go-ocpi/internal/cdr/v2.1.1"
	"github.com/satimoto/go-ocpi/test/mocks"
)

func TestCdrUnmarshal(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		dto := cdr.CdrDto{}
		response := []byte(`{
			"id": null,
			"start_date_time": null,
			"auth_id": null,
			"auth_method": null,
			"location": null,
			"currency": null,
			"tariffs": null,
			"charging_periods": null,
			"total_cost": null,
			"total_energy": null,
			"total_time": null,
			"last_updated": null
		}`)

		json.Unmarshal([]byte(`{}`), &dto)
		responseJson, _ := json.Marshal(dto)

		mocks.CompareJson(t, responseJson, response)
	})

	t.Run("Base data", func(t *testing.T) {
		dto := cdr.CdrDto{}
		request := []byte(`{
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
		}`)

		json.Unmarshal(request, &dto)
		responseJson, _ := json.Marshal(dto)

		mocks.CompareJson(t, responseJson, request)
	})

	t.Run("With charge points", func(t *testing.T) {
		dto := cdr.CdrDto{}
		request := []byte(`{
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
		}`)

		json.Unmarshal(request, &dto)
		responseJson, _ := json.Marshal(dto)

		mocks.CompareJson(t, responseJson, request)
	})
}
