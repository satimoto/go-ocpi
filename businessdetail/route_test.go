package businessdetail_test

import (
	"encoding/json"
	"testing"

	"github.com/satimoto/go-ocpi-api/businessdetail"
	"github.com/satimoto/go-ocpi-api/mocks"
)

func TestBusinessDetailUnmarshal(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		payload := businessdetail.BusinessDetailPayload{}
		response := []byte(`{
			"name": ""
		}`)

		json.Unmarshal([]byte(`{}`), &payload)
		responseJson, _ := json.Marshal(payload)

		mocks.CompareJson(t, responseJson, response)
	})

	t.Run("Name only", func(t *testing.T) {
		payload := businessdetail.BusinessDetailPayload{}
		request := []byte(`{
			"name": "Business Name"
		}`)

		json.Unmarshal(request, &payload)
		responseJson, _ := json.Marshal(payload)

		mocks.CompareJson(t, responseJson, request)
	})

	t.Run("Name and Website", func(t *testing.T) {
		payload := businessdetail.BusinessDetailPayload{}
		request := []byte(`{
			"name": "Business Name",
			"website": "https://business.com"
		}`)

		json.Unmarshal(request, &payload)
		responseJson, _ := json.Marshal(payload)

		mocks.CompareJson(t, responseJson, request)
	})

	t.Run("Name, Website and Logo", func(t *testing.T) {
		payload := businessdetail.BusinessDetailPayload{}
		request := []byte(`{
			"name": "Business Name",
			"website": "https://business.com",
			"logo": {
				"url": "https://business.com/logo.png",
				"category": "OPERATOR",
				"type": "png"
			}
		}`)

		json.Unmarshal(request, &payload)
		responseJson, _ := json.Marshal(payload)

		mocks.CompareJson(t, responseJson, request)
	})
}
