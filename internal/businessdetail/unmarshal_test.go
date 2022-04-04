package businessdetail_test

import (
	"encoding/json"
	"testing"

	"github.com/satimoto/go-ocpi-api/internal/businessdetail"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func TestBusinessDetailUnmarshal(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		dto := businessdetail.BusinessDetailDto{}
		response := []byte(`{
			"name": ""
		}`)

		json.Unmarshal([]byte(`{}`), &dto)
		responseJson, _ := json.Marshal(dto)

		mocks.CompareJson(t, responseJson, response)
	})

	t.Run("Name only", func(t *testing.T) {
		dto := businessdetail.BusinessDetailDto{}
		request := []byte(`{
			"name": "Business Name"
		}`)

		json.Unmarshal(request, &dto)
		responseJson, _ := json.Marshal(dto)

		mocks.CompareJson(t, responseJson, request)
	})

	t.Run("Name and Website", func(t *testing.T) {
		dto := businessdetail.BusinessDetailDto{}
		request := []byte(`{
			"name": "Business Name",
			"website": "https://business.com"
		}`)

		json.Unmarshal(request, &dto)
		responseJson, _ := json.Marshal(dto)

		mocks.CompareJson(t, responseJson, request)
	})

	t.Run("Name, Website and Logo", func(t *testing.T) {
		dto := businessdetail.BusinessDetailDto{}
		request := []byte(`{
			"name": "Business Name",
			"website": "https://business.com",
			"logo": {
				"url": "https://business.com/logo.png",
				"category": "OPERATOR",
				"type": "png"
			}
		}`)

		json.Unmarshal(request, &dto)
		responseJson, _ := json.Marshal(dto)

		mocks.CompareJson(t, responseJson, request)
	})
}
