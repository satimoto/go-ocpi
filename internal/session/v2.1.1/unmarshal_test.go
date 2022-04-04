package session_test

import (
	"encoding/json"
	"testing"

	session "github.com/satimoto/go-ocpi-api/internal/session/v2.1.1"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func TestSessionUnmarshal(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		dto := session.SessionDto{}
		response := []byte(`{
			"id": null,
			"start_datetime": null,
			"kwh": null,
			"auth_id": null,
			"auth_method": null,
			"location": null,
			"currency": null,
			"charging_periods": null,
			"status": null,
			"last_updated": null
		}`)

		json.Unmarshal([]byte(`{}`), &dto)
		responseJson, _ := json.Marshal(dto)

		mocks.CompareJson(t, responseJson, response)
	})

	t.Run("Base data", func(t *testing.T) {
		dto := session.SessionDto{}
		request := []byte(`{
			"id": "SESSION0001",
			"start_datetime": "2015-06-29T22:39:09Z",
			"kwh": 0,
			"auth_id": "DE8ACC12E46L89",
			"auth_method": "AUTH_REQUEST",
			"location": null,
			"currency": "EUR",
			"charging_periods": [],
			"status": "PENDING",
			"last_updated": "2015-06-29T22:39:09Z"
		}`)

		json.Unmarshal(request, &dto)
		responseJson, _ := json.Marshal(dto)

		mocks.CompareJson(t, responseJson, request)
	})

	t.Run("With charge points", func(t *testing.T) {
		dto := session.SessionDto{}
		request := []byte(`{
			"id": "SESSION0001",
			"start_datetime": "2015-06-29T22:39:09Z",
			"kwh": 15.342,
			"auth_id": "DE8ACC12E46L89",
			"auth_method": "AUTH_REQUEST",
			"location": null,
			"currency": "EUR",
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
			"status": "ACTIVE",
			"last_updated": "2015-06-29T22:39:09Z"
		}`)

		json.Unmarshal(request, &dto)
		responseJson, _ := json.Marshal(dto)

		mocks.CompareJson(t, responseJson, request)
	})
}
