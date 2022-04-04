package tariff_test

import (
	"encoding/json"
	"testing"

	tariff "github.com/satimoto/go-ocpi-api/internal/tariff/v2.1.1"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func TestTariffUnmarshal(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		payload := tariff.TariffPushPayload{}
		response := []byte(`{
			"id": null,
			"currency": null,
			"elements": null,
			"last_updated": null
		}`)

		json.Unmarshal([]byte(`{}`), &payload)
		responseJson, _ := json.Marshal(payload)

		mocks.CompareJson(t, responseJson, response)
	})

	t.Run("Base data", func(t *testing.T) {
		payload := tariff.TariffPushPayload{}
		request := []byte(`{
			"id": "TARIFF01",
			"currency": "EUR",
			"tariff_alt_url": "https://ev-power.de/",
			"elements": [],
			"last_updated": "2015-06-29T20:39:09Z"
		}`)

		json.Unmarshal(request, &payload)
		responseJson, _ := json.Marshal(payload)

		mocks.CompareJson(t, responseJson, request)
	})

	t.Run("With alt text and element", func(t *testing.T) {
		payload := tariff.TariffPushPayload{}
		request := []byte(`{
			"id": "TARIFF01",
			"currency": "EUR",
			"tariff_alt_text": [{
				"language": "en",
				"text": "2 euro p/hour"
			}],
			"tariff_alt_url": "https://ev-power.de/",
			"elements": [{
				"price_components": [{
					"type": "TIME",
					"price": 2.5,
					"step_size": 300
				}]
			}],
			"last_updated": "2015-06-29T20:39:09Z"
		}`)

		json.Unmarshal(request, &payload)
		responseJson, _ := json.Marshal(payload)

		mocks.CompareJson(t, responseJson, request)
	})

	t.Run("With multiple elements", func(t *testing.T) {
		payload := tariff.TariffPushPayload{}
		request := []byte(`{
			"id": "TARIFF01",
			"currency": "EUR",
			"tariff_alt_url": "https://ev-power.de/",
			"elements": [{
				"price_components": [{
					"type": "FLAT",
					"price": 2.5,
					"step_size": 1
				}]
			}, {
				"price_components": [{
					"type": "TIME",
					"price": 1,
					"step_size": 900
				}],
				"restrictions": {
					"max_power": 32
				}
			}, {
				"price_components": [{
					"type": "TIME",
					"price": 2,
					"step_size": 600
				}],
				"restrictions": {
					"min_power": 32,
					"day_of_week": ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY"]
				}
			}, {
				"price_components": [{
					"type": "TIME",
					"price": 1.25,
					"step_size": 600
				}],
				"restrictions": {
					"min_power": 32,
					"day_of_week": ["SATURDAY", "SUNDAY"]
				}
			}, {
				"price_components": [{
					"type": "PARKING_TIME",
					"price": 5,
					"step_size": 300
				}],
				"restrictions": {
					"start_time": "09:00",
					"end_time": "18:00",
					"day_of_week": ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY"]
				}
			}, {
				"price_components": [{
					"type": "PARKING_TIME",
					"price": 6,
					"step_size": 300
				}],
				"restrictions": {
					"start_time": "10:00",
					"end_time": "17:00",
					"day_of_week": ["SATURDAY"]
				}
			}],
			"last_updated": "2015-06-29T20:39:09Z"
		}`)

		json.Unmarshal(request, &payload)
		responseJson, _ := json.Marshal(payload)

		mocks.CompareJson(t, responseJson, request)
	})

	t.Run("With energy mix", func(t *testing.T) {
		payload := tariff.TariffPushPayload{}
		request := []byte(`{
			"id": "TARIFF01",
			"currency": "EUR",
			"tariff_alt_url": "https://ev-power.de/",
			"elements": [],
			"energy_mix": {
				"is_green_energy": false,
				"energy_sources": [{
					"source": "GENERAL_GREEN",
					"percentage": 35.9
				}, {
					"source": "GAS",
					"percentage": 6.3
				}, {
					"source": "COAL",
					"percentage": 33.2
				}, {
					"source": "GENERAL_FOSSIL",
					"percentage": 2.9
				}, {
					"source": "NUCLEAR",
					"percentage": 21.7
				}],
				"environ_impact": [{
					"source": "NUCLEAR_WASTE",
					"amount": 0.00006
				}, {
					"source": "CARBON_DIOXIDE",
					"amount": 372
				}],
				"supplier_name": "E.ON Energy Deutschland",
				"energy_product_name": "E.ON DirektStrom eco"
			},
			"last_updated": "2015-06-29T20:39:09Z"
		}`)

		json.Unmarshal(request, &payload)
		responseJson, _ := json.Marshal(payload)

		mocks.CompareJson(t, responseJson, request)
	})
}
